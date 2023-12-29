<script>
export default {
	info(): {
		return {
            deets:{
			errormsg: null,
            username: "",
            newUsername:"",
            userID:"",
            followerCount: 0,
            followingCount: 0,
            bannedCount:0,
            photoCount:0,
            },
		}
	},

	methods: {

		 async AccountInfo() {
            
			try {
                this.userID=localStorage.getItem("userID");
                this.username=localStorage.getItem("username")
				let response = await this.$axios.get("/user/${this.userID}/profile");
               
				this.username=response.info.username;
                this.userID=response.info.userID;

            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async setUser(){
        try{
            this.userID=localStorage.getItem("userID");
            let response=await this.$axios.put("/user/${this.userID}/setusername/${this.newUsername}")
            this.username=this.newUsername;
            localStorage.setItem("username",this.newUsername);
        }
        catch{
            this.errormsg=e.toString();
        }
        




    }


  
    },
}
</script>
<template>
    
    <h1> Account Information </h1>
    <input type="text" v-model="username" placeholder="set username">

    <button class="setuser" @click="setUser">update username</button>
    <div>
        <p> Username: {{ this.username }}</p>
    
    </div>
    
	
</template>

<style>
.setuser{
    display: flex;
    flex-direction:column;
    align-items:normal;
    justify-content:center;
    margin-top: 40px;
    


}



</style>