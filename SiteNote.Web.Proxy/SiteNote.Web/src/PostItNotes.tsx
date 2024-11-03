
import PostItNote from "./PostItNote";
import PlusIcon from "./assets/noteladplus.svg"
import "./styles/PostItNotes.css";



export default function PostItNotes() {



  return (
    <div className="PostItNotes WaveBackground">
      <PostItNote />
      <PostItNote />
      <div>
        <img className="plusIcon" src={PlusIcon}></img>
      </div>
    </div>

  );
}
