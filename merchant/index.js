import axios from "axios";
import express from "express";

const app = express();

app.use(express.json());

app.get("/ping", (req, res) => {
  res.send("hello world");
});

app.get("/get-merchant", async (req, res) => {
  try {
    const resp = await axios.get("http://localhost:9000/get-detail-toko");
    const detailToko = await resp.data
    const merchantDetail = {
        nama:"yanuar",
        nama_toko: detailToko.nama,
        jumlah_product:detailToko.jumlah_product
    }
    res.send({merhcant:merchantDetail});
  } catch (error) {
    console.log(error)
    res.send("asdad");
  }
});

app.listen("8000", () => {
  console.log("app running on port 8000");
});
