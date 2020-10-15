package handler

import (
	"net/http"

	"github.com/Bayu200720/mc/utils"
)

func AddMenu(w http.ResponseWriter, r *http.Request) {
	utils.WrapAPISuccess(w, r, "success", http.StatusOK)

}
