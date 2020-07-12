package response

import (
	"encoding/json"
	"net/http"

	"go.mongodb.org/mongo-driver/bson"
)

// JSON response
func JSON(w http.ResponseWriter, code int, msg string, payload interface{}) {
    w.WriteHeader(code)
    w.Header().Set("Content-Type", "application/json")
    if payload == false {
        json.NewEncoder(w).Encode(bson.M{"code": code, "success": false, "msg": msg, "data": nil})
        return
    }
    json.NewEncoder(w).Encode(bson.M{"code": code, "success": true, "msg": msg, "data": payload})
}