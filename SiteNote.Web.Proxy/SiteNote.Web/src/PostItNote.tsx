import { MutableRefObject, useEffect, useRef, useState } from "react";

import "./styles/PostItNote.css";
import Github, { GithubPlacement } from "@uiw/react-color-github";

export default function PostItNote() {
  let left: MutableRefObject<number> = useRef(10);
  let top: MutableRefObject<number> = useRef(10);
  const moving: MutableRefObject<boolean> = useRef(false);
  const hex = useRef("#ff");
  const [showPicker, setShowPicker] = useState(false);
  const inputFieldRef = useRef<HTMLTextAreaElement | null>(null);

  const [style, setStyle] = useState({
    left: left.current,
    top: top.current,
    "--color": hex.current,
  });

  useEffect(() => {
    // Add event listener on component mount
    window.addEventListener("mouseup", onMouseUp);

    // Add event listener on component mount
    window.addEventListener("mousemove", onMouseMove);

    // Remove event listener on component unmount
    return () => {
      window.removeEventListener("mouseup", onMouseUp);
      window.removeEventListener("mousemove", onMouseMove);
    };
  }, []); // Empty dependency array ensures this runs once on mount/unmount

  function onMouseDown() {
    moving.current = true;
  }

  function onMouseUp() {
    moving.current = false;
  }

  function onColorPickerPress() {
    console.log("HEJ");
    if (showPicker) {
      setShowPicker(false);
    } else {
      setShowPicker(true);
    }
  }

  function onMouseMove(e: MouseEvent) {
    if (moving.current) {
      left.current += e.movementX;
      top.current += e.movementY;
      setStyle({
        left: left.current,
        top: top.current,
        "--color": hex.current,
      });
    }
  }

  function handleColorChange(colorHex: string): void {

    hex.current = colorHex
    setStyle({
        left: left.current,
        top: top.current,
        '--color': hex.current
    })
  }

  function onPostItClick(){

        if(inputFieldRef.current){
            inputFieldRef.current.focus();
        }
        
    
    
  }

  return (
    <div className="PostIt draggable" onMouseDown={onMouseDown} onClick={onPostItClick} style={style}>
      <div className="overlay"></div>

      <div className="postItHeader">
        <div className="colorPickerArea">
          {showPicker ? (
            <Github
              color={hex.current}
              onChange={(color) => handleColorChange(color.hex)}
              placement={GithubPlacement.RightTop}
            />
          ) : (
            <></>
          )}
        </div>
        <button className="colorPicker" onClick={onColorPickerPress}>
          🖌️
        </button>
      </div>
      <div className="postItContent">
        <textarea ref={inputFieldRef}></textarea>
      </div>
    </div>
  );
}
