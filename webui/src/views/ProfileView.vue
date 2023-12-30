<script>
export default {
	data: function() {
		return {
            errormsg: null,
            userID: "",
            deets:{
			errormsg: null,
            username: "",
            newUsername:"",
            userID:"",
            followerCount: 2,
            followers:[],
            followingCount: 3,
            following:[],
            bannedCount:0,
            photoCount:0,

            },
		}
	},

    created(){
        this.AccountInfo();
    },

	methods: {

		 async AccountInfo() {
            
			try {

                console.log("grrr")
                this.userID=localStorage.getItem("userID");
                this.username=localStorage.getItem("username");

                console.log(this.userID);
                console.log(typeof this.userID);

				let response = await this.$axios.get("/username/"+ this.username + "/profile",);
                console.log(response);
               
				this.username=response.data.username;
                console.log(this.username);
                this.followerCount=response.data.followed_count;
                console.log(this.followerCount);
                this.followingCount=response.data.following_count;
                console.log(this.followerCount);
                this.photoCount=response.data.pic_numb

            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async getFollow(){
        try{
            console.log("method called")
            console.log(this.userID)

            let response=await this.$axios.get("/user/" + this.userID + "/followers")
            console.log(response);
        }
        catch{
            this.errormsg=e.toString();
        }
    },
    async following(){
        try{

            this.userID=localStorage.getItem("userID");
            let response=await this.$axios.get("/user/" + this.userID + "/followings")
        
        }
        catch{
            this.errormsg=e.toString();
        }
        




    },

    async mounted(){
         this.AccountInfo;
        await this.follow;
        await this.following;
        return;
    }



  
    },

} 

</script>
<template>
    
    <h1> Account Information </h1>

    <p class="username"> <b>Username:</b> <span> {{ this.username }}</span> </p>
    <div class="update">
    

    <input type="text" v-model="username" placeholder="set username">

    <button class="setuser" @click="AccountInfo">update username</button>

    </div>
    <div class="accountinfouser">
        
        <button class="followers" @click="getFollow" >followers: {{ this.followerCount }}</button>
        <button class="followers" >following: {{ this.followingCount }}</button>
        <button class="followers" >banned: {{ this.followingCount }}</button>
        <button class="followers" >pictures: {{ this.photoCount }}</button>

    
    </div>
    
	
</template>

<style>

.followers{
  
    border:0cqw;
    color: black;
    text-align: center;
    text-decoration: none;
    display: inline-block;
    font-size: 16px;
    cursor: pointer;
    border-radius: 8px;
    width: 160px;
    height: 24px;
    margin-top: 20px;
    margin-right: 50px;

}
.setuser{
    display: flex;
    flex-direction:column;
    align-items:normal;
    justify-content:center;
    margin-top: 10px;
    margin-right: 30px;


}

.accountinfouser{
    display: flex;
    flex-direction:column;
    justify-content:last baseline;
    margin-top: 50px;
    margin-right: 300px;


}

.username{
    display: flex;
    flex-direction:column;
    justify-content:last baseline;
    margin-top: 50px;
    margin-right: 300px;


}

input[type="text"] {
  padding: 5px;
  margin-right: 30px;
  margin-top: 10px;
}
.update{
    display: flex;
    flex-direction:row;
    align-items:normal;
    justify-content: last baseline;
    margin-top: 40px;
    


}


</style>
