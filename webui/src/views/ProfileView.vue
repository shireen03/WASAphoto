<script>
import LoadingSpinner from "../components/LoadingSpinner.vue";
export default {
    components: { LoadingSpinner },
	data: function() {
		return {
            errormsg: null,
           
            userUsername: localStorage.getItem("username"),
            newUser: "",
            userID: localStorage.getItem("userID"),
            followerCount: null,
            followingCount: null,
            bannedCount:null,
            photoCount:null,
            isFollow: false,
            isBan: false,
            isLoading: true,

            photos:{
                photo:[{
                photoID:0,
                likeCount:0,
                commentCount:0,
                date:"",
                file: "",
                isLike:false,
                comment: ""
            }],
            }

        
		}
	},
     created(){
        this.AccountInfo();
        
    },

	methods: {

      

        mounted(){
        this.AccountInfo()
    },

		 async AccountInfo() {
            
			try {
                console.log("idk");
                this.isLoading=true;
                this.userUsername=localStorage.getItem("username");
                this.userID=localStorage.getItem("userID");
                console.log(this.userUsername);
                await this.$axios.get("/username/"+ this.userUsername + "/profile").then((response) =>{
                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;
                    this.bannedCount=response.data.ban_count;
               

                });
                this.getPhotos();
                console.log("plsplspls work");

            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async refresh() {
		
			this.AccountInfo();
            
	},

    async updateUsername(){

        if(this.newUser==""){
            this.errormsg="you must enter a username!!!"

        }else{
            console.log(this.newUser);
            console.log(this.userID);
            let response=await this.$axios.put("/user/"+ this.userID + "/setusername/" + this.newUser);
            localStorage.setItem("username", response.data);
            console.log(response);
            this.refresh();
        }


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
                console.log("uploadddddd "+response);
                this.refresh();
        
            }
       
    
     catch{
            this.errormsg=e.toString();
        }


    },

    async getPhotos(){
        try{     
           
                let response=await this.$axios.get("/photo/upload/" + this.userID);
                console.log(response);
                this.photos=response.data;
                console.log(this.photos);

                for (let i=0;i<this.photos.photo.length();i++){
                    this.photos.photo[i].file= 'data:image/*;base64,' + this.photos.photo[i].file

                }
                
        
            }
       
    
     catch{
            this.errormsg=e.toString();
        }


    },


    
 

},
}

 

</script>
<template>
  
    
    <div class="titleButtons">

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

    </div>
    <div class="upload">
    <div class="btn-group btn-group-sm">
					<input type="file" accept="image/*" class="btn btn-xs" @change="uploadFile" ref="file">
					<button class="btn btn-default " @click="uploadPhoto">Upload</button>
				</div>
    </div>
   
    <div class="umidk">
    <div class="btn-group">
 
            <button class="btn"> followers: {{this.followerCount}} </button>
    
            <button class="btn"> following: {{this.followingCount}} </button>
            <button class="btn"> bans: {{this.bannedCount}} </button>

         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
    </div>
    </div>

    <div class="gro">
  <input type="text" id="username" v-model="this.newUser" placeholder="Enter new username" ><br>
  <button class="btn btn-outline-secondary" id="submit" @click="updateUsername"> submit  </button>
  </div>







    

  


	
</template>

<style>



.titleButtons{
    display:flex;
    flex-direction:row;
    align-items:normal;
    justify-content:first baseline;
    margin-top: 37px;
    gap: 80px;

}





.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

.photo{
    display:grid;
    align-items:first baseline;
}






</style>
