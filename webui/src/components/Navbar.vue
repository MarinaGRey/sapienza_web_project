<script>
export default {
  data(){
    return{
      textVar: "",
    }
  },
  methods:{
    logout(){
      localStorage.removeItem('token')
      this.$emit('logoutNavbar',false)
    },
    goBackHome(){
      this.$emit('requestUpdateView',"/home")
    },
    searchFunc(){
      this.$emit('searchNavbar',this.textVar)
      this.textVar=""
    },
    myProfile(){
      this.$emit('requestUpdateView',"/users/"+localStorage.getItem('token'))
    },
  },
}
</script>
<template>
  <nav class="navbar navbar-expand-lg bg-light d-flex justify-content-between sticky-top mb-3 my-nav bg-transparent">
    <div class="col-4">
      <a class="navbar-brand ms-2 d-flex" @click="goBackHome" style="margin-right: 10px;">
        <div class="fw-bold">WASAPhoto</div>
      </a>
    </div>

    <!-- -->
    <div class="col-4">
      <form class="form-inline my-2 my-lg-0 d-flex justify-content-center m-auto">
        <input class="form-control mr-sm-2 w-50" v-model="textVar" type="search" placeholder="Search users">
        <button class="btn btn-light ms-2" type="submit" @click.prevent="searchFunc">Search</button>
      </form>
    </div>

    <div class="col-4 d-flex justify-content-end">
      <button @click="myProfile" type="button" class="btn btn-light me-2">
        Profile
      </button>

      <button @click="logout" type="button" class="btn btn-light">
        Logout
      </button>
    </div>
  </nav>
</template>

<style>
  .navbar-brand div {
    font-weight: bold;
  }
</style>
