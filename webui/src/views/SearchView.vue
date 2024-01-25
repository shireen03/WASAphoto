<script>
export default {
	data: function() {
		return {
			errormsg: null,
      searchUsername: "",
      banStatus:false,
      userStatus: false,
      userID: "",
      search: "",
      username:"",
      
		}
	},
	methods: {

    async searchUser(){
		try {
            this.username=localStorage.getItem("username");
            this.userID=localStorage.getItem("userID");

            let response=await this.$axios.get("/username/"+ this.search + "/ban/" + this.username, {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.banStatus=response.data

            let response2=await this.$axios.get("/username/"+ this.search + "/checkUser", {
                headers: {
                    Authorization: "Bearer " + localStorage.getItem("userID")
                }
            });
            this.userStatus=response2.data

            if(this.search==""){
                this.errormsg = "search cannot be empty"
            }
            else if(this.userStatus==false){
              this.errormsg="user doesnt exist"
            }else if(this.search==this.username){
                this.errormsg="Why you searching yourself"
            }else if(this.banStatus==true){
              this.errormsg="you're banned by this user."
            }else{                
              this.$router.push({path: "user/" + this.search + "/account"});
            }
        } 

        catch (e) {
				this.errormsg = e.toString();
		}
		
    },
	
    },
}
</script>

<template>
  		<br>
			<h4>Search User</h4>
		<br><br>
    
      <div class="group">
          <input class="text" type="text" style="height: 40px; width: 400px" v-model="search" placeholder="Search User">
          <button class="btn btn-outline-secondary" id="submit" style="height: 40px;" @click="searchUser" > search  </button>
      </div>
        <ErrorMsg v-if="errormsg" :msg="errormsg"></ErrorMsg>


		
</template>
<style>

.group{
  display: flex;
  align-items: center;
  justify-content: center;
}

</style>
