package endpoints

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func getUser(w http.ResponseWriter, r *http.Request) {
	log.Println("getuser called")
	vars := mux.Vars(r)
	user, ok := vars["USERNAME"]
	if !ok {
		http.Error(w, "Cannot find username in request", http.StatusBadRequest)
		return
	}
	if _, ok := users[user]; ok {
		sendJSONResponse(w,
			struct {
				Username    string `json:"username"`
				Description string `json:"description"`
			}{user, "有天良，有底线的一般般的人。甘于平凡甚至平庸，凡事不为最先，不耻最后，随着大流走。有时也会标新立异，但是骨子里还是传统守旧。十分忠诚敬业，恪守职业道德。做什么，爱什么，做一行，精一行，不求闻达，但求无愧。以为这世上只有不体面的人，没有不体面的职业。能够理解朋友的优点，能够接受朋友的缺点，任何时候不会为了一己私利，背信弃义，出尔反尔。不会有什么了不起的作为，俯仰无愧天地。做不到的话，过头的话尽量不说，说过的话尽量要做到。不违心地吹捧任何人，也不会违心地委屈自己。严格律己，宽以待人。不把自己的任何观念强加给任何人。会不满现状，但总是为了自我完善。写完上述文字，自我表白是不是太多了，是我必须检点的。能够屏蔽所有的不愉快，余生的每一天都是上苍的馈赠，因此我会热爱生活，充满感恩。六十年回望，优点不会没有，缺点不会很多。余生展望，一如既往，任何人都不能够突然地就君子豹变成另外一个完全不同于自己的人。自我认识是完全无法与认识你这个个体所存在的社会截然分开的，知人论世往往说的是别人，其实同样适用于你自己。汲黯卧治，非匹夫之能事；老聃知足，实晚景之余晖。"})
		return
	}
	http.Error(w, "Cannot find user", http.StatusNotFound)
}
