import { useState } from "react";
import axios from "axios";

const ImageUpload = () => {
  const [imageURL, setImageURL] = useState("");
  console.log(imageURL);

  const handleUpload = async (e: any) => {
    e.preventDefault();

    const formData = new FormData();
    formData.append("image", e.target.image.files[0]);

    try {
      const response = await axios.post(
        "http://localhost:4200/api/upload",
        formData,
        {
          headers: {
            "Content-Type": "multipart/form-data",
          },
        }
      );
      console.log(response);

      setImageURL(response.data.url);
    } catch (error) {
      console.error(error);
    }
  };

  return (
    <div>
      <form onSubmit={handleUpload}>
        <input type="file" name="image" />
        <button type="submit">Upload</button>
      </form>
      {imageURL && (
        <div>
          <p>Uploaded image:</p>
          <img src={imageURL} alt="Uploaded" />
        </div>
      )}
    </div>
  );
};

export default ImageUpload;
