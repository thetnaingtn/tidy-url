import React, { useRef, useEffect } from 'react';
import rough from 'roughjs/bin/rough';

const RoughInputButton: React.FC = () => {
  const canvasRef = useRef<HTMLCanvasElement | null>(null);

  useEffect(() => {
    const canvas = canvasRef.current;
    if (!canvas) return;

    const rc = rough.canvas(canvas);
    const ctx = canvas.getContext('2d');
    if (!ctx) return;

    // Clear previous drawings
    ctx.clearRect(0, 0, canvas.width, canvas.height);

    // Rough rectangle for input
    rc.rectangle(20, 20, 200, 40, {
      roughness: 1.2,
      bowing: 1,
      stroke: '#333',
      strokeWidth: 1.5,
    });

    // Rough rectangle for button
    rc.rectangle(240, 20, 100, 40, {
      roughness: 1.2,
      bowing: 1,
      stroke: '#333',
      strokeWidth: 1.5,
      fill: '#fdf6e3',
      fillStyle: 'hachure',
    });
  }, []);

  return (
    <div style={{ position: 'relative', width: 360, height: 80 }}>
      <canvas
        ref={canvasRef}
        width={360}
        height={80}
        style={{
          position: 'absolute',
          top: 0,
          left: 0,
          zIndex: 0,
        }}
      />
      <input
        placeholder="Type here..."
        style={{
          position: 'absolute',
          top: 20,
          left: 20,
          width: 200,
          height: 40,
          border: 'none',
          background: 'transparent',
          paddingLeft: 10,
          fontSize: '16px',
          zIndex: 1,
        }}
      />
      <button
        style={{
          position: 'absolute',
          top: 20,
          left: 240,
          width: 100,
          height: 40,
          border: 'none',
          background: 'transparent',
          cursor: 'pointer',
          zIndex: 1,
          fontSize: '16px',
        }}
      >
        Submit
      </button>
    </div>
  );
};

export default RoughInputButton;
