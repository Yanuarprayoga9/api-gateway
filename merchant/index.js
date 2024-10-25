import express from "express"


const app = express()

app.use(express.json())

app.get("/ping",(req,res)=>{
    res.send("hello world")
})
app.listen("8000",()=>{
    console.log("app running on port 8000")
})