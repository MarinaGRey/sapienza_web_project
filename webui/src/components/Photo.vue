<script>
export default {
	data(){
		return{
			liked: false,
			allComments: [],
			allLikes: [],
		}
	},

	props: ["owner","username","likes","comments","upload_date","photo_id","file","isOwner"], 

	methods:{
		async loadPhoto(){

			try {
        		let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/photos/" + this.photo_id + "/file", {responseType: "blob"})
        		this.photoURL = URL.createObjectURL(response.data)
			} catch (e) {
				this.errormsg = e.toString()
			}

		},


		async deletePhoto(){
			try{
				// Delete photo: /users/:id/photos/:photo_id
				await this.$axios.delete("/users/"+localStorage.getItem('token')+"/photos/"+this.photo_id)
				// location.reload()
				this.$emit("removePhoto",this.photo_id)
			}catch(e){
				//
			}
		},

		photoOwnerClick: function(){
			this.$router.replace("/users/"+this.owner)
		},

		async toggleLike() {

			if(this.isOwner){ 
				return
			}

			const bearer = localStorage.getItem('token')

			try{
				if (!this.liked){

					// Put like: /users/:id/photos/:photo_id/likes"
					await this.$axios.post("/users/"+ localStorage.getItem('token') +"/photos/"+this.photo_id+"/likes")
					this.allLikes.push({
						user_id: localStorage.getItem('token'),
						nickname: bearer
					})

				}else{
					// Delete like: /users/:id/photos/:photo_id/likes"
					await this.$axios.delete("/users/"+ localStorage.getItem('token')  +"/photos/"+ this.photo_id +"/likes")
					this.allLikes.pop()
				}

				this.liked = !this.liked;
			}catch(e){
				//
			}
      		
    	},

		removeCommentFromList(value){
			this.allComments = this.allComments.filter(item=> item.comment_id !== value)
		},

		addCommentToList(comment){
			this.allComments.push(comment)
		},
	},
	
	async mounted(){
		await this.loadPhoto()

		if (this.likes != null){
			this.allLikes = this.likes
		}

		if (this.likes != null){
			this.liked = this.allLikes.some(obj => obj.userid === localStorage.getItem('token'))
		}
		if (this.comments != null){
			this.allComments = this.comments
		}
		
		
	},

}
</script>

<template>
  <div class="container-fluid mt-3 mb-5">

    <LikeModal :modal_id="'like_modal'+photo_id" :likes="allLikes" />

    <CommentModal :modal_id="'comment_modal'+photo_id" :comments_list="allComments" :photo_owner="owner" :photo_id="photo_id"
      @eliminateComment="removeCommentFromList" @addComment="addCommentToList" />

    <div class="d-flex flex-row justify-content-center">

      <div class="card my-card">
        <div class="d-flex justify-content-end">

          <button v-if="isOwner" :class="['my-btn', 'my-delete-btn', 'me-2']" @click="deletePhoto">
            Delete
          </button>

        </div>
        <div class="d-flex justify-content-center photo-background-color">
          <img :src="photoURL" class="card-img-top img-fluid">
        </div>

        <div class="card-body">

          <div class="container">

            <div class="d-flex flex-row justify-content-end align-items-center mb-2">

              <button class="my-btn m-0 p-1 me-auto" @click="photoOwnerClick">
                From {{ username }}
              </button>

              <button :class="['my-like-btn', liked ? 'not-liked-bg' : 'liked-bg']" @click="toggleLike">
                {{ liked ? 'Unlike' : 'Like' }} {{ allLikes.length }}
              </button>

              <button class="my-btn m-0 p-1" data-bs-toggle="modal" :data-bs-target="'#comment_modal' + photo_id">
                Comments {{ allComments != null ? allComments.length : 0 }}
              </button>
            </div>

            <div class="d-flex flex-row justify-content-start align-items-center ">
              <p> Uploaded on {{ upload_date }}</p>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
  .photo-background-color {
    background-color: grey;
  }

  .my-card {
    width: 27rem;
    border-color: black;
    border-width: thin;
  }

  .my-btn {
    background-color: transparent;
    border: none;
    padding: 0;
    text-decoration: none;
    color: black;
    transition: color 0.3s ease;
  }

  .my-like-btn {
    background-color: transparent;
    border: none;
    padding: 8px 16px;
    text-decoration: none;
    cursor: pointer;
    border-radius: 5px;
    transition: background-color 0.3s ease, color 0.3s ease;
  }

  .liked-bg {
    background-color: #4CAF50; /* Green */
  }

  .not-liked-bg {
    background-color: #f44336; /* Red */
  }

  .my-delete-btn {
    background-color: #f44336; /* Red */
    color: white;
    padding: 8px 16px;
    font-size: 12px;
  }
</style>
