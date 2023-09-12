package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"

	// Importing the general-purpose Cosmos blockchain client
	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
	"github.com/ignite/cli/ignite/pkg/cosmosclient"

	// Importing the types package of your blog blockchain
	"blog/x/blog/types"
)

var (
	accountName = "alice"
)

func createPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("cosmos"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	body := r.FormValue("body")

	err = createPost(ctx, client, account, title, body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post created successfully")
}

func createPost(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, title string, body string) error {
	addr, err := account.Address("cosmos")
	if err != nil {
		log.Fatal(err)
	}

	msg := &types.MsgCreatePost{
		Creator: addr,
		Title:   title,
		Body:    body,
	}
	_, err = client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func updatePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("cosmos"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Get account from the keyring
	account, err := client.Account(accountName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	title := r.FormValue("title")
	body := r.FormValue("body")
	strId := r.FormValue("id")

	intId, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid uint64 value", http.StatusBadRequest)
		return
	}

	err = updatePost(ctx, client, account, title, body, intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post updated successfully")
}

func updatePost(ctx context.Context, client cosmosclient.Client,
	account cosmosaccount.Account, title string, body string, id uint64) error {

	addr, err := account.Address("cosmos")
	if err != nil {
		log.Fatal(err)
	}

	msg := &types.MsgUpdatePost{
		Creator: addr,
		Id:      id,
		Title:   title,
		Body:    body,
	}

	_, err = client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func deletePostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("cosmos"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	account, err := client.Account(accountName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	strId := r.FormValue("id")
	intId, err := strconv.ParseUint(strId, 10, 64)
	if err != nil {
		http.Error(w, "Invalid uint64 value", http.StatusBadRequest)
		return
	}

	err = deletePost(ctx, client, account, intId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintln(w, "Post deleted successfully")
}

func deletePost(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, id uint64) error {
	addr, err := account.Address("cosmos")
	if err != nil {
		log.Fatal(err)
	}

	msg := &types.MsgDeletePost{
		Creator: addr,
		Id:      id,
	}

	_, err = client.BroadcastTx(ctx, account, msg)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func listPostHandler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	// Create a Cosmos client instance
	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix("cosmos"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	// Call listPost() to retrieve the posts
	queryResp := listPost(ctx, client)
	// Serialize the response to JSON and send it in the response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(queryResp)
}

func listPost(ctx context.Context, client cosmosclient.Client) *types.QueryAllPostResponse {
	queryClient := types.NewQueryClient(client.Context())
	// Query the blockchain using the client's `PostAll` method
	// to get all posts store all posts in queryResp
	queryResp, err := queryClient.PostAll(ctx, &types.QueryAllPostRequest{})
	if err != nil {
		log.Fatal(err)
	}
	return queryResp
}

func main() {
	r := mux.NewRouter()

	// Define API route for CRUD functions
	r.HandleFunc("/posts/list", listPostHandler)
	r.HandleFunc("/posts/create", createPostHandler)
	r.HandleFunc("/posts/update", updatePostHandler)
	r.HandleFunc("/posts/delete", deletePostHandler)

	fmt.Println("Server is listening on :8080")
	log.Fatal(http.ListenAndServe(":8080",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
			handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		)(r)))

}

// splitaaaaaaaaaaaaaaaaaaaaaaaaaaaaa

// // package main

// // import (
// // 	"context"
// // 	"fmt"
// // 	"log"
// // 	"net/http"

// // 	// Importing the general purpose Cosmos blockchain client
// // 	"github.com/ignite/cli/ignite/pkg/cosmosaccount"
// // 	"github.com/ignite/cli/ignite/pkg/cosmosclient"

// // 	// Importing the types package of your blog blockchain
// // 	"blog/x/blog/types"
// // )

// // func listPost(ctx2 context.Context, client cosmosclient.Client) {

// // 	queryClient := types.NewQueryClient(client.Context())

// // 	// Query the blockchain using the client's `PostAll` method
// // 	// to get all posts store all posts in queryResp
// // 	queryResp, err := queryClient.PostAll(ctx2, &types.QueryAllPostRequest{})
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	// Print response from querying all the posts
// // 	fmt.Println("\n\nAll Posts:\n\n")
// // 	fmt.Println(queryResp)
// // }

// // func createPost(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, title string, body string) {
// // 	addr, err := account.Address("cosmos")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	msg := &types.MsgCreatePost{
// // 		Creator: addr,
// // 		Title:   title,
// // 		Body:    body,
// // 	}

// // 	_, err = client.BroadcastTx(ctx, account, msg)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// // func updatePost(ctx context.Context, client cosmosclient.Client,
// // 	account cosmosaccount.Account, title string, body string, id uint64) {

// // 	addr, err := account.Address("cosmos")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	msg := &types.MsgUpdatePost{
// // 		Creator: addr,
// // 		Id:      id,
// // 		Title:   title,
// // 		Body:    body,
// // 	}

// // 	_, err = client.BroadcastTx(ctx, account, msg)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// // func deletePost(ctx context.Context, client cosmosclient.Client, account cosmosaccount.Account, id uint64) {
// // 	addr, err := account.Address("cosmos")
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	msg := &types.MsgDeletePost{
// // 		Creator: addr,
// // 		Id:      id,
// // 	}

// // 	_, err = client.BroadcastTx(ctx, account, msg)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}
// // }

// // func main() {

// // 	ctx := context.Background()
// // 	addressPrefix := "cosmos"

// // 	// Create a Cosmos client instance
// // 	client, err := cosmosclient.New(ctx, cosmosclient.WithAddressPrefix(addressPrefix))
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	//init account
// // 	accountName := "alice"
// // 	// Get account from the keyring
// // 	account, err := client.Account(accountName)
// // 	if err != nil {
// // 		log.Fatal(err)
// // 	}

// // 	fmt.Println("Server is listening on :8080")
// // 	http.ListenAndServe(":8080", nil)

// // 	createPost(ctx, client, account, "title2", "body2")
// // 	updatePost(ctx, client, account, "updateTitle", "updateBody", 1)
// // 	deletePost(ctx, client, account, 18)
// // 	listPost(ctx, client)

// // }
