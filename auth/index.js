import express from "express"
import dotenv from "dotenv"
dotenv.config()

const user = [
    {
        username:"yanuar",
        password:"pass"
    },
    {
        username:"ahmad",
        password:"pass"
    },
]
const {PORT} = process.env

const app = express();

app.use(express.json())

app.get("/ping",async(req,res)=>{
    res.send("ping")
})
app.post("/login",async(req,res)=>{
    
})
app.listen(PORT,()=>{
    console.log(`app running on port : ${PORT}`)
})