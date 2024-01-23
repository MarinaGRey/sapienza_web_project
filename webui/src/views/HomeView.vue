<script>
export default {
	data: function () {
		return {
			errormsg: null,
			photos: [],
		}
	},

	methods: {
		
		async loadStream() {
			try {
				this.errormsg = null
				// Home get: "/users/:id/homescreen"
				let response = await this.$axios.get("/users/" + localStorage.getItem('token') + "/homescreen")

				if (response.data != null){
					this.photos = response.data
				}
				
			} catch (e) {
				this.errormsg = e.toString()
			}
		}
	},

	async mounted() {
		await this.loadStream()
	}

}
</script>

<template>
	<div class="container-fluid">
		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>

		<div class="row">
			<Photo
				v-for="(photo,index) in photos"
				:key="index"
				:owner="photo.userid"
				:username="photo.username"
				:photo_id="photo.photoid"
				:file="photo.file"
				:comments="photo.comments != nil ? photo.comments : []"
				:likes="photo.likes != nil ? photo.likes : []"
				:upload_date="photo.date"
			/>
		</div>

		<div v-if="photos.length === 0" class="row ">
			<h1 class="d-flex justify-content-center mt-5" style="color: white;">There's no content yet, follow somebody!</h1>
		</div>


	</div>
</template>

<style>
</style>
