package main

import (
	"encoding/json"
	"log"
	"net/http"

	"k8s.io/api/admission/v1beta1"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/serializer"
)

const (
	certFile = "/etc/webhook/certs/cert.pem"
	keyFile  = "/etc/webhook/certs/key.pem"

	//certFile = "./server.pem"
	//keyFile  = "./server-key.pem"
)

var (
	runtimeScheme = runtime.NewScheme()
	codecs        = serializer.NewCodecFactory(runtimeScheme)
	deserializer  = codecs.UniversalDeserializer()
	defaulter     = runtime.ObjectDefaulter(runtimeScheme)
)

func main() {
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.Println("start self web hook ..")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("hello\n"))
	})
	http.HandleFunc("/mutate", webhookMutate)
	http.HandleFunc("/validate", webhookValidate)
	err := http.ListenAndServeTLS(":443", certFile, keyFile, nil)
	if err != nil {
		panic(err)
	}
}

func webhookMutate(w http.ResponseWriter, r *http.Request) {
	var review v1beta1.AdmissionReview
	defer func() {
		err := json.NewEncoder(w).Encode(review)
		if err != nil {
			log.Println(err)
			return
		}
	}()
	err := json.NewDecoder(r.Body).Decode(&review)
	log.Println("at webhookMutatet ", review)
	if err != nil {
		review.Response = &v1beta1.AdmissionResponse{Result: &metav1.Status{Message: err.Error()}}
		return
	}

	req := review.Request
	switch req.Kind.Kind {
	case "Deployment":
		var deployment appsv1.Deployment
		err := json.Unmarshal(req.Object.Raw, &deployment)
		if err != nil {
			log.Println(err)
		}
		log.Println(deployment)
	}
	//str := base64.StdEncoding.EncodeToString([]byte(`[{"op": "add", "path": "/spec/replicas", "value": 1}]`))
	pt := v1beta1.PatchTypeJSONPatch
	review.Response = &v1beta1.AdmissionResponse{
		Allowed:   true,
		PatchType: &pt,
		Patch:     []byte(`[{"op": "add", "path": "/spec/replicas", "value": 1}]`),
	}
}

func webhookValidate(w http.ResponseWriter, r *http.Request) {
	var review v1beta1.AdmissionReview
	defer func() {
		err := json.NewEncoder(w).Encode(review)
		if err != nil {
			log.Println(err)
			return
		}
	}()
	err := json.NewDecoder(r.Body).Decode(&review)
	log.Println("at webhookValidate ", review)
	if err != nil {
		review.Response = &v1beta1.AdmissionResponse{Result: &metav1.Status{Message: err.Error()}}
		return
	}

	req := review.Request
	switch req.Kind.Kind {

	}
	log.Println("===========")
	log.Println(string(req.Object.Raw))
	review.Response = &v1beta1.AdmissionResponse{Allowed: true}
}
