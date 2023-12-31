<script>
export default {
	data: function() {
		return {
            errormsg: null,
            deets:{
			errormsg: null,
            userUsername: "",
            userID:"",
            followerCount: 2,
            followingCount: 3,
            bannedCount:0,
            photoCount:0,
            isFollow: false,
            isBan: false,

            },
		}
	},


    created(){
        this.AccountInfo();
    },



	methods: {

       

		 async AccountInfo() {
            
			try {
                this.userUsername=localStorage.getItem("username");
                this.userID=localStorage.getItem("userID");
                console.log(this.userUsername);
                let response = await this.$axios.get("/username/"+ this.userUsername + "/profile");
                console.log(response);
                this.followerCount=response.data.followed_count;
                console.log(this.followerCount);
                this.followingCount=response.data.following_count;
                console.log(this.followerCount);
                this.photoCount=response.data.pic_numb;

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
    async refresh() {
			this.loading = true;
			this.errormsg = null;
			try {
				let response = await this.$axios.get("/username/"+ this.userUsername + "/profile");
				console.log(response);
                console.log("refreshed");

			} catch (e) {
				this.errormsg = e.toString();
			}
			this.loading = false;
		},
	},
	
    async getFollow(){
        try{     
           
            
          
            
        }
       
    
     catch{
            this.errormsg=e.toString();
        }


    },

    async uploadPhoto(){
        try{     
           
                let response=await this.$axios.post("/photo/upload/" + this.userID);
                console.log(response);
                this.refresh();
        
            }
       
    
     catch{
            this.errormsg=e.toString();
        }


    },
    async mounted(){
        this.AccountInfo();
        await this.follow;
        await this.following;
        this.refresh();
        return;
    },

    }

 

</script>
<template>
    <head>
        <script> AccountInfo()</script>
    </head>
    
    <div class="titleButtons">

    <h2> <span>{{this.userUsername}}</span> Account   </h2>

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

    <div class="upload">
    <div class="btn-group btn-group-sm">
					<input type="file" accept="image/*" class="btn btn-xs" @change="uploadFile" ref="file">
					<button class="btn btn-default " @click="uploadPhoto">Upload</button>
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
.upload{
    text-align: 100px;
    margin-left: 600px;
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
