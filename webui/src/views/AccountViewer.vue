<script>
export default {
	data: function() {
		return {
            errormsg: null,
            userID: "",
            currentUser: "",
            deets:{
			errormsg: null,
            userUsername: "",
            userID:"",
            followerCount: 2,
            followingCount: 3,
            bannedCount:0,
            photoCount:0,
            isFollow: true,

            },
		}
	},

    created(){
        this.CheckBanFollow();
    },

	methods: {

        async CheckBanFollow(){
            try {
                this.userUsername=this.$route.params.username;
                this.currentUser=localStorage.getItem("username");
                console.log(this.currentUser);

                let response = await this.$axios.get("/username/"+ this.currentUser + "/follow/" + this.userUsername);
                console.log(response)

                
            } catch (error) {
                
            }

        },

		 async AccountInfo() {
            
			try {
                this.userUsername=this.$route.params.username;
                console.log(this.userUsername);
                let response = await this.$axios.get("/username/"+ this.userUsername + "/profile");
                console.log(response);
                this.followerCount=response.data.followed_count;
                console.log(this.followerCount);
                this.followingCount=response.data.following_count;
                console.log(this.followerCount);
                this.photoCount=response.data.pic_numb

                /*
                console.log("plsplspls work")

                this.userUsername=this.$route.params.username;
                console.log(this.userUsername);


                console.log(response)
                this.userID=localStorage.getItem("userID");
                this.username=localStorage.getItem("username");

                console.log(this.userID);
                console.log(typeof this.userID);

               
				this.username=response.data.username;
                console.log(this.username);
                
*/
            
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
    
    <div class="titleButtons">

    <h2> <span>{{this.userUsername}}</span> Account   </h2>

    <div class="buttons">
        <button class="followers" @click="getFollow" >{{isFollowing==true? "follow":"unfollow" }}</button>
        <button class="followers" @click="getFollow">ban</button>
    </div>
    </div>
  
    <div class="counts">
        <div class="followCount">
            <h6> followers </h6>
             <h6> <span>{{this.followerCount}}</span></h6>
        </div>

        <div class="followingCount">
            <h6> following </h6>
            <h6><span>{{ this.followingCount }}</span></h6>
        </div>

        <div class="photoCount">
            <h6> photos </h6>
            <h6><span>{{ this.photoCount }}</span></h6>
        </div>
    </div>


	
</template>

<style>

.titleButtons{
    display:-webkit-flex;
    flex-direction:row;
    align-items:normal;
    justify-content:first baseline;
    margin-top: 37px;
    gap: 80px;

}
.followers{
    margin-top: 2px;
}



.followers{
    width: 100px;
    height: 40px;
}
.counts{
    display: flex;
    flex-direction:row;
    align-items:normal;
    justify-content:baseline;
    margin-top: 50px;
    gap: 70px;

}
.followCount{
    display: flex;
    flex-direction:row;
    align-items:normal;
    justify-content:baseline;
    gap: 10px;
}
.followingCount{
    display: flex;
    flex-direction:row;
    align-items:normal;
    justify-content:baseline;
    gap: 10px;
}
.photoCount{
    display: flex;
    flex-direction:row;
    align-items:normal;
    justify-content:baseline;
    gap: 10px;
}




</style>
