<script>
export default {
	data: function() {
		return {
			errormsg: null,
      username: "",
      userID: "",
		}
	},
	methods: {

	async dologin(){
		try {
				
				console.log("Login method called")


				let response = await this.$axios.post("/session",{
                    username: this.username
                });
        
        console.log(response);
      
				localStorage.setItem("username", this.username);
        localStorage.setItem("userID", response.data);

        this.userID = response.data;
       
        this.$axios.defaults.headers.common['Authorization']=`Bearer ${localStorage.getItem("userID")}`;

        console.log("uhh");
        console.log(localStorage.getItem("userID"));




        this.$router.push({path: "/session"});
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
	



	},
  
}
</script>

<template>
    <div class="center">
	  <h1 class="h1">Login</h1>
    <div class="gro">
  <input type="text" id="username" v-model="this.username" placeholder="Enter new username" ><br>
  <button class="btn btn-outline-dark" id="submit" @click="dologin" > submit  </button>
  </div>

</div>
  </template>

<style>

.center {
  display: flex;
  flex-direction:column;
  align-items: center;
  justify-content: center;
  min-height: 50vh;

}
.h1 {
  display: flex;
  flex-direction:row;
  align-items: center;
  justify-content: center;


}


.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}



</style>

