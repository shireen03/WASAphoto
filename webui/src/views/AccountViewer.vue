<script>
export default {
	data: function() {
		return {
            errormsg: null,
            currentUserID: "",
            currentUser: "",
            
			errormsg: null,
            userUsername: "",
            userIDUser:"",
            followerCount: 0,
            followingCount: 0,
            bannedCount:0,
            photoCount:0,
            isFollow: false,
            isBan: false,
            checkban:false,

          
		}
	},


    created(){
        this.CheckBanFollow()
    },



	methods: {

        async CheckBanFollow(){
            try {
                this.userUsername=this.$route.params.username;
                this.currentUser=localStorage.getItem("username");
                console.log("aight")
                let checkBan = await this.$axios.get("/username/"+ this.userUsername + "/ban/" + this.currentUser);
                console.log(checkBan);
                this.checkban=checkBan.data;
                if(this.checkban){
                    this.$router.push({path: "/Error"});

                    
                }else{
                
                console.log(this.currentUser);

                let response = await this.$axios.get("/username/"+ this.currentUser + "/follow/" + this.userUsername);
                console.log(response);
                console.log(response.data);
                this.isFollow=response.data;
                if(this.isFollow){
                    document.getElementById("follow").innerHTML = "unfollow";
                    this.AccountInfo(); 
                }else{
                    document.getElementById("follow").innerHTML = "follow"; 
                    this.AccountInfo();
                }

                let banResponse = await this.$axios.get("/username/"+ this.currentUser + "/ban/" + this.userUsername);
                console.log(banResponse);
                console.log("ban: "+ banResponse.data);
                this.isBan=banResponse.data;
                if(this.isBan){
                    document.getElementById("ban").innerHTML = "unban"; 
                }else{
                    document.getElementById("ban").innerHTML = "ban"; 
                }
                this.refresh();
       

            }

                
            } catch (error) {
                
            }

        },

		 async AccountInfo() {
            
			try {
                console.log("wtf");
                await this.$axios.get("/username/"+ this.userUsername + "/profile").then((response) =>{
                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;

                });
                console.log(response);
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
           
            console.log("follow method called");

            this.currentUser=localStorage.getItem("username");
            console.log(this.isFollow);
            if(this.isFollow){
                let response=await this.$axios.delete("/username/" + this.currentUser + "/follow/" + this.userUsername);
                this.isFollow=false;
                document.getElementById("follow").innerHTML = "follow"; 
                 console.log(response);
                 this.refresh();
            }else{
                console.log("ban check:" + this.isBan);
                if(this.isBan==false){
                let response=await this.$axios.post("/username/" + this.currentUser + "/follow/" + this.userUsername);
                this.isFollow=true;
                console.log(response);
                document.getElementById("follow").innerHTML = "unfollow";
                this.refresh();
                }
            


          
            
        }
       
    }
     catch{
            this.errormsg=e.toString();
        }


    },

    async refresh() {
        console.log("refreshhh");
		
        this.AccountInfo();
        
    },

    async getBan(){
        try{     
           
            console.log("ban method called");

            this.currentUser=localStorage.getItem("username");
            console.log(this.isBan);
            if(this.isBan){
                let response=await this.$axios.delete("/username/" + this.currentUser + "/ban/" + this.userUsername);
                this.isBan=false;
                document.getElementById("ban").innerHTML = "ban"; 
                 console.log(response);
                 this.refresh();
            }else{
                let response=await this.$axios.post("/username/" + this.currentUser + "/ban/" + this.userUsername);

                this.isBan=true;
                console.log(response);
                document.getElementById("ban").innerHTML = "unban";
                

                console.log(this.isFollow);
                if(this.isFollow){
                    let response=await this.$axios.delete("/username/" + this.currentUser + "/follow/" + this.userUsername);
                    let wtf=await this.$axios.delete("/username/" + this.userUsername + "/follow/" + this.currentUser);
                    this.isFollow=false;
                    document.getElementById("follow").innerHTML = "follow"; 

                    
                }
                this.refresh();
            
        }
       
    }
     catch{
            this.errormsg=e.toString();
        }


    },

    },}

 

</script>
<template>
  
    
    <div class="titleButtons">

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

    </div>
    <div class="upload">
    <div class="btn-group">
        <button class="btn btn-outline-dark " id="follow" @click="getFollow">wtv </button>

		<button class="btn btn-outline-dark " id="ban" @click="getBan">wtv</button>
				</div>
    </div>
    
    <div class="btn-group">
 
            <button class="btn"> followers: {{this.followerCount}} </button>
    
            <button class="btn"> following: {{this.followingCount}} </button>
         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
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




.btn-group{
  background-color: rgb(213, 213, 213);
  text-align: start;

}





</style>
