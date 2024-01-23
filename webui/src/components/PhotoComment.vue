<script>
export default {
    data(){
        return {
            user: "",
        }
    },
	props: ['content','author','photo_owner','comment_id','photo_id','nickname'],

    methods:{
        async deleteComment(){
            try{
                // Delete comment: "/users/:id/photos/:photo_id/comments/:comment_id"
                await this.$axios.delete("/users/"+localStorage.getItem('token')+"/photos/"+this.photo_id+"/comments/"+this.comment_id)

                this.$emit('eliminateComment',this.comment_id)

            }catch (e){
                console.log(e.toString())
            }
        },
    },

    mounted(){
        this.user = localStorage.getItem('token')
    }

}
</script>

<template>
    <div class="container-fluid">

        <hr>
        <div class="row">
            <div class="col-10">
                <h5>{{nickname}}</h5>
            </div>

            <div class="col-2">
                <button v-if="user === photo_owner || user === author" class="btn my-delete-btn" @click="deleteComment">
                    Delete
                </button>
            </div>

        </div>

        <div class="row">
            <div class="col-12">
                {{content}}
            </div>

        </div>
        <hr>
    </div>
</template>

<style>
/**/

</style>