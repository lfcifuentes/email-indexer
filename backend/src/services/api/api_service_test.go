package api

import (
	"encoding/base64"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// TestGet verifica que el método Get realiza correctamente una petición GET sin autenticación.
func TestGet(t *testing.T) {
	// Crea un servidor de prueba que responde con un JSON.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("Esperado GET, recibido %s", r.Method)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"message": "hello"}`))
	}))
	defer ts.Close()

	apiSvc := NewApiService("", "")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := apiSvc.Get(ts.URL, headers)
	if err != nil {
		t.Fatalf("Error inesperado en Get: %v", err)
	}

	expected := `{"message": "hello"}`
	if strings.TrimSpace(string(resp)) != expected {
		t.Errorf("Respuesta inesperada: esperado %s, obtenido %s", expected, string(resp))
	}
}

// TestPostWithAuth verifica que una petición POST con autenticación envía correctamente las cabeceras.
func TestPostWithAuth(t *testing.T) {
	expectedBody := map[string]interface{}{
		"query": "search term",
	}
	// Crea un servidor de prueba que leerá el body y verificará la autenticación.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Verifica Basic Auth.
		username, password, ok := r.BasicAuth()
		if !ok || username != "myUser" || password != "myPass" {
			t.Errorf("Autenticación básica incorrecta, recibido username:%s, password:%s", username, password)
		}

		// Verifica el método.
		if r.Method != http.MethodPost {
			t.Errorf("Esperado POST, recibido %s", r.Method)
		}

		// Decodifica el body y verifícalo.
		var body map[string]interface{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			t.Errorf("Error al decodificar el body: %v", err)
		}

		if body["query"] != expectedBody["query"] {
			t.Errorf("Body inesperado, esperado %v, obtenido %v", expectedBody, body)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status": "ok"}`))
	}))
	defer ts.Close()

	apiSvc := NewApiService("myUser", "myPass")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := apiSvc.PostWithAuth(ts.URL, headers, expectedBody)
	if err != nil {
		t.Fatalf("Error inesperado en PostWithAuth: %v", err)
	}

	expectedResp := `{"status": "ok"}`
	if strings.TrimSpace(string(resp)) != expectedResp {
		t.Errorf("Respuesta inesperada: esperado %s, obtenido %s", expectedResp, string(resp))
	}
}

// TestGetWithAuth verifica que el método GetWithAuth envía la cabecera de autenticación correcta.
func TestGetWithAuth(t *testing.T) {
	// Crea un servidor de prueba que verifique la cabecera Authorization.
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		// Construimos el valor esperado para Basic Auth.
		expectedAuth := "Basic " + base64.StdEncoding.EncodeToString([]byte("myUser:myPass"))
		if authHeader != expectedAuth {
			t.Errorf("Authorization incorrecto, esperado %s, obtenido %s", expectedAuth, authHeader)
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"authorized": true}`))
	}))
	defer ts.Close()

	apiSvc := NewApiService("myUser", "myPass")
	headers := map[string]string{
		"Content-Type": "application/json",
	}
	resp, err := apiSvc.GetWithAuth(ts.URL, headers)
	if err != nil {
		t.Fatalf("Error inesperado en GetWithAuth: %v", err)
	}

	expected := `{"authorized": true}`
	if strings.TrimSpace(string(resp)) != expected {
		t.Errorf("Respuesta inesperada: esperado %s, obtenido %s", expected, string(resp))
	}
}
