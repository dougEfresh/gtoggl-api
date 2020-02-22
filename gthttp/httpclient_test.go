package gthttp

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type TestLogger struct {
	Testing *testing.T
}

func (l *TestLogger) Printf(format string, v ...interface{}) {
	l.Testing.Logf(format, v...)
}

var mockRequest = struct {
	path, query       string // request
	contenttype, body string // response
}{
	path:        DefaultUrl + "/projects",
	contenttype: "application/json",
	body:        "{ \"data\": {\"nothing\": true}}",
}
var handler = func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", mockRequest.contenttype)
	if strings.Contains(r.URL.Path, "/sessions") {
		w.Header().Set("Set-Cookie", "__Host-timer-session=MTM2MzA4MJa8jA3OHxEdi1CQkFFQ180SUFBUkFCRUFBQVlQLUNBQUVHYzNSeWFXNW5EQXdBQ25ObGMzTnBiMjVmYVdRR2MzUnlhVzVuREQ0QVBIUnZaMmRzTFdGd2FTMXpaWE56YVc5dUxUSXRaalU1WmpaalpEUTVOV1ZsTVRoaE1UaGhaalpqWkRkbU5XWTJNV0psWVRnd09EWmlPVEV3WkE9PXweAkG7kI6NBG-iqvhNn1MSDhkz2Pz_UYTzdBvZjCaA==; Path=/; Expires=Wed, 13 Mar 2013 09:54:38 UTC; Max-Age=86400; HttpOnly")
		io.WriteString(w, "")
	} else {
		io.WriteString(w, mockRequest.body)
	}
}

func TestHandleGet(t *testing.T) {

	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	testLogger := &TestLogger{t}
	tc, err := NewClient("token", SetURL(server.URL), SetErrorLogger(testLogger))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	raw, err := tc.GetRequest(server.URL + "/any")
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	var tResp TogglResponse
	err = json.Unmarshal(*raw, &tResp)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}

	if tResp.Data == nil {
		t.Fatalf("Get: %v", err)
	}
}

func TestHandlePost(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	testLogger := &TestLogger{t}
	tc, err := NewClient("token", SetURL(server.URL), SetErrorLogger(testLogger))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	raw, err := tc.PostRequest(server.URL+"/any", &TogglResponse{Data: nil})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	var tResp TogglResponse
	err = json.Unmarshal(*raw, &tResp)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if tResp.Data == nil {
		t.Fatalf("Get: %v", err)
	}
}

func TestHandlePut(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	testLogger := &TestLogger{t}
	tc, err := NewClient("token", SetURL(server.URL), SetErrorLogger(testLogger))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	raw, err := tc.PutRequest(server.URL+"/any", &TogglResponse{Data: nil})
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	var tResp TogglResponse
	err = json.Unmarshal(*raw, &tResp)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if tResp.Data == nil {
		t.Fatalf("Get: %v", err)
	}
}

func TestHandleDelete(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()
	testLogger := &TestLogger{t}
	tc, err := NewClient("token", SetURL(server.URL), SetErrorLogger(testLogger))
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	raw, err := tc.DeleteRequest(server.URL+"/any", nil)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	var tResp TogglResponse
	err = json.Unmarshal(*raw, &tResp)
	if err != nil {
		t.Fatalf("Get: %v", err)
	}
	if tResp.Data == nil {
		t.Fatalf("Get: %v", err)
	}
}

func TestClientDefaults(t *testing.T) {
	client, err := NewClient("")
	if err == nil {
		t.Fatal("Error should have been thrown. No Token given")
	}
	server := httptest.NewServer(http.HandlerFunc(handler))
	defer server.Close()

	httpClient := &http.Client{}
	testLogger := &TestLogger{t}
	client, err = NewClient("abc1234567890def", SetURL(server.URL), SetErrorLogger(testLogger), SetTraceLogger(testLogger), SetInfoLogger(testLogger), SetHttpClient(httpClient))
	if err != nil {
		t.Fatal(err)
	}
	if client.Url != server.URL {
		t.Error("Url not blah; " + client.Url)
	}
	if client.traceLog != testLogger || client.errorLog != testLogger || client.infoLog != testLogger {
		t.Error("Loggers not set ")
	}
	if client.cookie == nil {
		t.Errorf("Session Cookie not found not defined\n")
	}
}

func Test400(t *testing.T) {
	var h = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mockRequest.contenttype)
		w.WriteHeader(400)
	}
	server := httptest.NewServer(http.HandlerFunc(h))
	defer server.Close()
	testLogger := &TestLogger{t}
	client, err := NewClient("abc1234567890def", SetURL(server.URL), SetTraceLogger(testLogger))
	fmt.Printf("%s\n", err)
	if err == nil {
		t.Fatal("Should be 400")
	}

	if client != nil {
		t.Fatal("client should be null")
	}

}

func Test404(t *testing.T) {
	var h = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", mockRequest.contenttype)
		if strings.Contains(r.URL.Path, "/sessions") {
			w.Header().Set("Set-Cookie", "__Host-timer-session=MTM2MzA4MJa8jA3OHxEdi1CQkFFQ180SUFBUkFCRUFBQVlQLUNBQUVHYzNSeWFXNW5EQXdBQ25ObGMzTnBiMjVmYVdRR2MzUnlhVzVuREQ0QVBIUnZaMmRzTFdGd2FTMXpaWE56YVc5dUxUSXRaalU1WmpaalpEUTVOV1ZsTVRoaE1UaGhaalpqWkRkbU5XWTJNV0psWVRnd09EWmlPVEV3WkE9PXweAkG7kI6NBG-iqvhNn1MSDhkz2Pz_UYTzdBvZjCaA==; Path=/; Expires=Wed, 13 Mar 2013 09:54:38 UTC; Max-Age=86400; HttpOnly")
			io.WriteString(w, "")
		} else {
			w.WriteHeader(404)
		}
	}
	server := httptest.NewServer(http.HandlerFunc(h))
	defer server.Close()
	testLogger := &TestLogger{t}
	client, err := NewClient("abc1234567890def", SetURL(server.URL), SetTraceLogger(testLogger))

	if err != nil {
		t.Fatal("Error")
	}
	r, err := client.GetRequest(server.URL + "/whatever")
	if r != nil {
		t.Fatal("r should be null")
	}
	if err != nil {
		t.Fatalf("err should be null %v", err)
	}

}
