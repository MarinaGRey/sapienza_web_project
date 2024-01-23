<script>
export default {
	data: function() {
		return {
			users: [],
			errormsg: null,
		}
	},

	props:['searchValue'],

	watch:{
		searchValue: function(){
			this.loadSearchedUsers()
		},
	},

	methods:{
		async loadSearchedUsers(){
			this.errormsg = null;
			if (
				this.searchValue === undefined ||
				this.searchValue === "" || 
				this.searchValue.includes("?") ||
				this.searchValue.includes("_")){
				this.users = []
				return 
			}

			try {
				// Search user (PUT):  "/users"
				let response = await this.$axios.put("/users",{
						username: this.searchValue.trim(),
				});
				// this.users = response.data
				if (response.data.userid != "") {
					this.$router.replace("/users/"+response.data.userid)
				}

			} catch (e) {
				 this.errormsg = "There is no user with that name";
			}
		},

	},

	async mounted(){
		// Check if the user is logged
		if (!localStorage.getItem('token')){
			this.$router.replace("/login")
		}
		await this.loadSearchedUsers()
		
	},
}
</script>

<template>
	<div>
		<p v-if="users.length == 0" class="no-result-text d-flex justify-content-center"> No users found.</p>

		<ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
	</div>
</template>

<style>
/**/
</style>
