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
            photos:[{
                photoID:0,
                username:"",
                like_count:0,
                comment_count:0,
                date:"",
                photo: "",
                isLike:false,
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
        this.CheckBanFollow()
    },



	methods: {

        async CheckBanFollow(){
            try {
                this.currentUserID=localStorage.getItem("userID");

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
                let response=await this.$axios.get("/username/"+ this.userUsername + "/profile");
                    this.userIDUser=response.data.userID;

                    this.followerCount=response.data.followed_count;
                    this.followingCount=response.data.following_count;
                    this.photoCount=response.data.pic_numb;

               
                console.log(response.data);
                console.log("brovski "+ this.userIDUser);
               this.getPhotos();
            
			} catch (e) {
				this.errormsg = e.toString();
			}
		
    },
          
    async getPhotos(){
           
           let response=await this.$axios.get("/user/" + this.currentUserID + "/photo/" + this.userIDUser);
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


    },    async LikePhoto(photoID){

    console.log("putting like")

    let response=await this.$axios.post("/user/" + this.currentUserID + "/photo/" + photoID +"/like");
    console.log(response.data);
    this.getPhotos();
      
       

    },

    async unLikePhoto(photoID){

    console.log("deleting like")
    let response=await this.$axios.delete("/user/" + this.currentUserID + "/photo/" + photoID +"/like");
    console.log(response.data);

    this.getPhotos();



    },

    },




    }

 

</script>
<template>  

<head><meta name="viewport" content="width=device-width, initial-scale=1">
    <link rel="stylesheet" href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"></head>



    
    <div class="titleButtons">

    <h2 > <span>{{this.userUsername}}</span> Account   </h2>

    </div>

    <div class="btn-group">
        <button class="btn btn-outline-dark " id="follow" @click="getFollow">wtv </button>

		<button class="btn btn-outline-dark " id="ban" @click="getBan">wtv</button>
	</div>
    
    <div class="btn-group">
 
            <button class="btn"> followers: {{this.followerCount}} </button>
    
            <button class="btn"> following: {{this.followingCount}} </button>
         
   
            <button class="btn"> photos: {{this.photoCount}} </button>
    </div>


    <br>

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
                    <h4>{{ photo.foll }}</h4>
    
                <button class="fa fa-heart"  @click="this.unLikePhoto(photo.photoID)" v-if="photo.isLiked==true"> {{ photo.like_count }}</button>
                <button class="fa fa-hearto" @click="this.LikePhoto(photo.photoID)" v-if="photo.isLiked==false" > {{ photo.like_count }}</button>

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



.titleButtons{
    display:-webkit-flex;
    flex-direction:row;
    align-items:normal;
    justify-content:first baseline;
    margin-top: 37px;
    gap: 80px;

}


.btn-group{
  background-color: rgb(213, 213, 213);
  text-align: start;

}

.row{

    margin-top: 37px;
}

.card{
    box-shadow: 0 4px 8px 0 rgba(0,0,0,0.2);
    width: 100px;
    height:800px;
    

}


.container {
  margin-top: 10px;
  margin-right: 50px;
}


.gro{
    display:flex;
    flex-direction: row;
    align-items:flex-end;
    justify-content:start;
}

.fa-hearto {
  color: rgb(255, 255, 255);
  background-color: grey;
  cursor: pointer;
  height: 27px;
}
.fa-heart {
  color: grey;
  cursor: pointer;
  background-color: black;

  height: 27px;
}














</style>
