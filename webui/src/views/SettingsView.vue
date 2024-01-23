<script>
export default {
	data: function () {
		return {
			errormsg: null,
			nickname: "",
		}
	},

	methods:{
		async modifyNickname(){
			try{
				// Nickname put: /users/:id
				let resp = await this.$axios.put("/users/"+this.$route.params.id,{
					username: this.nickname,
				})

				this.nickname=""
			}catch (e){
				this.errormsg = e.toString();
			}
		},
	},

}
</script>

<template>
  <div class="container-fluid">
    <div class="row">
      <div class="col d-flex justify-content-center mb-2">
        <h1>Settings</h1>
      </div>
    </div>

    <div class="row mt-2">
      <div class="col d-flex flex-column align-items-center">
        <p class="mb-2 font-size-lg">Change username:</p>
        <div class="input-group mb-3 w-25">
          <input
            type="text"
            class="form-control w-25"
            placeholder="Write your new nickname"
            maxlength="16"
            minlength="3"
            v-model="nickname"
          />
          <div class="input-group-append">
            <button
              class="btn btn-outline-secondary"
              @click="modifyNickname"
              :disabled="nickname === null || nickname.length > 16 || nickname.length < 3 || nickname.trim().length === 0"
            >
              Modify
            </button>
          </div>
        </div>
        <p class="text-muted font-size-sm">Remember that the name must be between 3 and 16 characters long.</p>
      </div>
    </div>

    <div class="row">
      <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>
    </div>
  </div>
</template>

<style>
  .font-size-lg {
    font-size: 1.5rem;
  }

  .font-size-sm {
    font-size: 1.2rem; 
  }
</style>

