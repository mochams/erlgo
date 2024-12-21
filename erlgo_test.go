package erlgo

import (
	"bytes"
	"encoding/binary"
	"io"
	"os"
	"testing"
)

// mockStdin helps test ReadFromErlang by providing test input
type mockStdin struct {
	*bytes.Reader
}

func (m *mockStdin) Close() error { return nil }

// mockStdout helps test WriteToErlang by capturing output
type mockStdout struct {
	*bytes.Buffer
}

func (m *mockStdout) Sync() error { return nil }

func TestReadFromErlang(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		want    []byte
		wantErr bool
	}{
		{
			name:    "simple message",
			input:   buildTestMessage([]byte("hello")),
			want:    []byte("hello"),
			wantErr: false,
		},
		{
			name:    "empty message",
			input:   buildTestMessage([]byte{}),
			want:    []byte{},
			wantErr: false,
		},
		{
			name:    "incomplete length",
			input:   []byte{0, 0},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "incomplete message",
			input:   append([]byte{0, 0, 0, 5}, []byte("hi")...),
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace stdin with our mock
			oldStdin := os.Stdin
			r, w, _ := os.Pipe()
			os.Stdin = r

			// Write test input
			w.Write(tt.input)
			w.Close()

			// Run test
			got, err := ReadFromErlang()

			// Restore stdin
			os.Stdin = oldStdin

			// Check results
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadFromErlang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && !bytes.Equal(got, tt.want) {
				t.Errorf("ReadFromErlang() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestWriteToErlang(t *testing.T) {
	tests := []struct {
		name    string
		input   []byte
		wantErr bool
	}{
		{
			name:    "simple message",
			input:   []byte("hello"),
			wantErr: false,
		},
		{
			name:    "empty message",
			input:   []byte{},
			wantErr: false,
		},
		{
			name:    "large message",
			input:   bytes.Repeat([]byte("a"), 1024),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Replace stdout with our mock
			oldStdout := os.Stdout
			r, w, _ := os.Pipe()
			os.Stdout = w

			// Run test
			err := WriteToErlang(tt.input)

			// Close write end and read result
			w.Close()
			var buf bytes.Buffer
			io.Copy(&buf, r)

			// Restore stdout
			os.Stdout = oldStdout

			// Check results
			if (err != nil) != tt.wantErr {
				t.Errorf("WriteToErlang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !tt.wantErr {
				// Verify length prefix
				got := buf.Bytes()
				if len(got) < 4 {
					t.Error("WriteToErlang() wrote too few bytes")
					return
				}

				length := binary.BigEndian.Uint32(got[:4])
				if int(length) != len(tt.input) {
					t.Errorf("WriteToErlang() wrote length %d, want %d", length, len(tt.input))
				}

				// Verify message content
				if !bytes.Equal(got[4:], tt.input) {
					t.Errorf("WriteToErlang() wrote message %v, want %v", got[4:], tt.input)
				}
			}
		})
	}
}

// Helper function to build test messages with length prefix
func buildTestMessage(content []byte) []byte {
	length := make([]byte, 4)
	binary.BigEndian.PutUint32(length, uint32(len(content)))
	return append(length, content...)
}
