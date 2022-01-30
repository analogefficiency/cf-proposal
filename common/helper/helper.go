package helper

import (
	"bytes"
	"cf-proposal/common/logservice"
	"cf-proposal/common/types"
	"cf-proposal/domain/model"
	"encoding/hex"
	"encoding/json"
	"hash/fnv"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

type CustomError struct {
	Message string
}

func (ce *CustomError) Error() string {
	return ce.Message
}

func GetShortUrl(longUrl string) string {

	hash := fnv.New32a()
	hash.Write([]byte(longUrl))
	return hex.EncodeToString(hash.Sum(nil))
}

func GetEncodedPayload(v interface{}) *bytes.Buffer {
	b := new(bytes.Buffer)
	json.NewEncoder(b).Encode(v)
	return b
}

func GetTestResponseBodyJson(v interface{}, r *httptest.ResponseRecorder) (interface{}, error) {
	encodedBody, err := ioutil.ReadAll(r.Result().Body)
	defer r.Result().Body.Close()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(encodedBody, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

func HandleHttpError(w http.ResponseWriter, r *http.Request, err error, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	response, _ := json.Marshal(model.ErrorDto{
		Error: err.Error(),
	})

	w.Write(response)
	logservice.LogError(http.StatusBadRequest, r.Method, types.Path(r.URL.Path), err)
}

func HandleHttpOk(w http.ResponseWriter, r *http.Request, v interface{}, code int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	jsonResp, _ := json.Marshal(v)
	w.Write(jsonResp)
	logservice.LogHttpRequest(code, r.Method, types.Path(r.URL.Path))
}
