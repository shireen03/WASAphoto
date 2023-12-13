<script>
export default {
	data: function() {
		return {
			errormsg: null,
            username: "",
            followerCount: 0,
            followerArray: [],
            followingCount: 0,
            followingArray:[],
            banArray:[],
            photoArray:[],
            photoCount:0,
		}
	},

	methods: {

		 async AccountInfo() {
            
			try {
				console.log("Login method called");

				let response = await this.$axios.get("/user",{
                    username: this.username
                });
				this.username = response.data

				localStorage.setItem("username", this.username);


                this.$router.push({path: '/session'});
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async setUser(){


    }


  
    },
}
</script>
<template>
	<h1> Account info </h1>
    <input type="text" v-model="username" placeholder="set username">

    <button class="setuser" @click="setUser">update username</button>
    <div class ="followercount">
    <h6>Followers: {{ followerCount }} </h6>


    </div>
	
</template>

<style>
.followercount{
    display: flex;
    flex-direction:column;
    align-items:normal;
    justify-content:baseline;
    margin-top: 40px;
    


}



</style>