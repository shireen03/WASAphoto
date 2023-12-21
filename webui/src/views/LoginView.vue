<script>
export default {
	data: function() {
		return {
			errormsg: null,
      username: "",
      userID:0,

		}
	},
	methods: {

	async dologin(){
		try {
				
				console.log("Login method called");

				let response = await this.$axios.post("/session",{
                    username: this.username
                });
				this.username = response.data.username
        this.userID = response.data.userID

				localStorage.setItem("username", this.username);
        localStorage.setItem("userID", this.userID);



        this.$router.push({path: '/session'});
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
	  <div class="input-row">
      	<input type="text" v-model="username" placeholder="Enter Username">
      	<button class="log" @click="dologin">Login</button>
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
.input-row {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
}

input[type="text"] {
  padding: 5px;
  margin-right: 30px;
  margin-top: 40px;
}

.log{
  padding:5px 10px;
  margin-top: 40px;
}
</style>

