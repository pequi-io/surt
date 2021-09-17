package s3service

import (
	"bytes"
	"context"
	"io/ioutil"
	"strconv"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type mockGetObjectAPI func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)

func (m mockGetObjectAPI) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	return m(ctx, params, optFns...)
}

func TestGetObjectBody(t *testing.T) {
	cases := []struct {
		client func(t *testing.T) S3GetObject
		bucket string
		key    string
		expect []byte
	}{
		{
			client: func(t *testing.T) S3GetObject {
				return mockGetObjectAPI(func(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
					t.Helper()
					if params.Bucket == nil {
						t.Fatal("expect bucket to not be nil")
					}
					if e, a := "surt-bucket", *params.Bucket; e != a {
						t.Errorf("expect %v, got %v", e, a)
					}
					if params.Key == nil {
						t.Fatal("expect key to not be nil")
					}
					if e, a := "surt-file", *params.Key; e != a {
						t.Errorf("expect %v, got %v", e, a)
					}

					return &s3.GetObjectOutput{
						Body: ioutil.NopCloser(bytes.NewReader([]byte("Hello world"))),
					}, nil
				})
			},
			bucket: "surt-bucket",
			key:    "surt-file",
			expect: []byte(""),
		},
	}

	for i, tt := range cases {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			ctx := context.TODO()
			content, err := GetObjectBody(ctx, tt.client(t), tt.bucket, tt.key)
			if err != nil {
				t.Fatalf("expect no error, got %v", err)
			}
			if e, a := tt.expect, content; !bytes.Equal(e, a) {
				t.Errorf("expect %v, got %v", e, a)
			}
		})
	}
}
