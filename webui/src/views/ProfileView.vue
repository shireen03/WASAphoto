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
            likeID: 0,

        
            photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLiked:true,
                comment: "",
                comments:[{
                        commentID:0,
                        commentUser: "",
                        comment:"",

                    }],
                }
            ],
            }

        
		},
     created(){

        this.AccountInfo();
        //this.getPhotos();
        
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
                if(this.photoCount>0){
                    this.getPhotos();
                };
                
                console.log("plsplspls work");

            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
    async refresh() {
		
			this.AccountInfo();
            
	},


    async LikePhoto(photoID){

console.log("putting like")

        let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/like");
    console.log(response.data);
    this.getPhotos();
    
    

},
async unLikePhoto(photoID){

console.log("deleting like")
let response=await this.$axios.delete("/user/" + this.userID + "/photo/" + photoID +"/like");
console.log(response.data);

this.getPhotos();



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
            console.log("hahaha")
                // let pic=document.getElementById("inputImage").files[0];
                // console.log(pic);
                // const read= new FileReader();

                // read.readAsDataURL(pic);
                // console.log("sheesh")

                const uploadpic=this.$refs.pico.files[0];
                // read.onload= async () =>{};
                let response= await this.$axios.post("/photo/upload/" + this.userID, uploadpic);


                this.AccountInfo();
                
                
                this.refresh();
        
            }
       
    
     catch(e){
            this.errormsg=e.toString();
        }


    },

       
    async getPhotos(){
           
                let response=await this.$axios.get("/photo/upload/" + this.userID);
                console.log(response.data);
                this.photos=response.data;
                console.log(this.photos);
                console.log(this.photos[0].photo)
                console.log(this.photos.length);
                for (let i=0;i<this.photos.length;i++){
                    
                    this.photos[i].photo= 'data:image/png;base64,'+ this.photos[i].photo;
                    console.log("omg   "+this.photos[i].photo);
                    console.log("slay   "+this.photos[i].photoID);
                    console.log("slay   "+this.photos[i].likeCount);
                      
                }
                
        
    },

    async uploadComment(photoID,yuhcomment){
        console.log(yuhcomment);
           
           let response=await this.$axios.post("/user/" + this.userID + "/photo/" + photoID +"/comment", { comment: yuhcomment });
           console.log(response.data);
          
           
   
},




    
 

},
}

 

</script>
<template>
    <head><meta name="viewport" content="width=device-width, initial-scale=1">
        <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>
  
    

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

  
    <div class="upload">
    <div class="btn-group btn-group-sm">
					<input type="file" accept="image/*" class="btn btn-xs" id="inputImage" ref="pico" @change="uploadFile" >

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
    <button class="btn btn-outline-dark" id="submit" @click="updateUsername"> submit  </button>
    </div>

    <br>

    <head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap@4.6.2/dist/css/bootstrap.min.css">
 
    </head>

    <div class="container">
        <div class="card-columns">
            <div v-for="photo in this.photos" style="width=300px">
                <div class="card">
                    <img class="card-img-top" :src=photo.photo  alt="unavailable" >
                    <hr>
                    <div class="card-body">
        
                    <button class="fa fa-heart" v-if="photo.isLiked==true" @click="this.unLikePhoto(photo.photoID)"> {{ photo.like_count }}</button>
                    <button class="fa fa-hearto"  v-if="photo.isLiked==false" @click="this.LikePhoto(photo.photoID)"> {{ photo.like_count }}</button>

                    <button  type="button" @click="uploadComment( photo.photoID, photo.comment)">comments</button>

                    <br>
                    <br>
                    <p > likes: {{ photo.like_count }}<br>
                    comments: {{ photo.comment_count }}<br>
                     date: {{ photo.date }}</p>
   
                    <div class="gro">
                    <input type="text" id="comment" v-model="photo.comment" class="form-control" placeholder="input comment">
                        <div class="input-group-append">
                            <button class="btn btn-outline-dark" type="button" @click="uploadComment(photo.photoID, photo.comment)">post</button>
                        </div>
                    </div>
                    </div>
                
                </div>

            </div>
        </div>
     </div>
     


</template>


<style>


.fa-hearto {
  color: red;
  background-color: grey;
  cursor: pointer;
  height: 27px;
}
.fa-heart {
  color: red;
  cursor: pointer;
  background-color: red;
  height: 27px;
}




.commenter{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:flex-start;
 
}
.card{
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    width: 100px;
    height:800px;
    border-color: blue;
    

}



.combtn{
    border-color: black;
    margin-right: 200px;
    width: 50px;
    height:30px;
    

}

.container {
  margin-top: -100px;
  margin-right: 50px;
}






.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}






</style>
