/**
 * Whiteboard - 무한 캔버스 그림판
 */
class Whiteboard {
    constructor(canvas) {
        this.canvas = canvas;
        this.ctx = canvas.getContext('2d');

        // 캔버스 크기를 윈도우에 맞춤
        this.resizeCanvas();

        // 상태
        this.isDrawing = false;
        this.isPanning = false;
        this.lastX = 0;
        this.lastY = 0;

        // 변환 상태
        this.offsetX = 0;
        this.offsetY = 0;
        this.scale = 1;

        // 그리기 설정
        this.penSize = 4;
        this.penColor = '#000000';

        // 그리기 히스토리 (무한 캔버스를 위해)
        this.strokes = [];
        this.currentStroke = null;

        // 이벤트 바인딩
        this.bindEvents();

        // 초기 렌더링
        this.render();
    }

    resizeCanvas() {
        this.canvas.width = window.innerWidth;
        this.canvas.height = window.innerHeight;
    }

    bindEvents() {
        // 마우스 이벤트
        this.canvas.addEventListener('mousedown', this.handleMouseDown.bind(this));
        this.canvas.addEventListener('mousemove', this.handleMouseMove.bind(this));
        this.canvas.addEventListener('mouseup', this.handleMouseUp.bind(this));
        this.canvas.addEventListener('wheel', this.handleWheel.bind(this), { passive: false });

        // 터치 이벤트
        this.canvas.addEventListener('touchstart', this.handleTouchStart.bind(this));
        this.canvas.addEventListener('touchmove', this.handleTouchMove.bind(this));
        this.canvas.addEventListener('touchend', this.handleTouchEnd.bind(this));

        // 윈도우 리사이즈
        window.addEventListener('resize', () => {
            this.resizeCanvas();
            this.render();
        });

        // 컨텍스트 메뉴 방지 (우클릭)
        this.canvas.addEventListener('contextmenu', (e) => e.preventDefault());
    }

    // 화면 좌표 -> 캔버스 좌표 변환
    screenToCanvas(screenX, screenY) {
        return {
            x: (screenX - this.offsetX) / this.scale,
            y: (screenY - this.offsetY) / this.scale
        };
    }

    // 캔버스 좌표 -> 화면 좌표 변환
    canvasToScreen(canvasX, canvasY) {
        return {
            x: canvasX * this.scale + this.offsetX,
            y: canvasY * this.scale + this.offsetY
        };
    }

    handleMouseDown(e) {
        const isRightClick = e.button === 2;
        const isSpacePressed = e.shiftKey || isRightClick;

        if (isSpacePressed) {
            // 팬 모드
            this.isPanning = true;
            this.canvas.classList.add('panning');
            this.lastX = e.clientX;
            this.lastY = e.clientY;
        } else {
            // 그리기 모드
            this.isDrawing = true;
            const pos = this.screenToCanvas(e.clientX, e.clientY);
            this.startDrawing(pos.x, pos.y);
        }
    }

    handleMouseMove(e) {
        // 좌표 표시 업데이트
        const pos = this.screenToCanvas(e.clientX, e.clientY);
        document.getElementById('coords').textContent =
            `X: ${Math.round(pos.x)}, Y: ${Math.round(pos.y)}`;

        if (this.isPanning) {
            const dx = e.clientX - this.lastX;
            const dy = e.clientY - this.lastY;
            this.pan(dx, dy);
            this.lastX = e.clientX;
            this.lastY = e.clientY;
        } else if (this.isDrawing) {
            const pos = this.screenToCanvas(e.clientX, e.clientY);
            this.draw(pos.x, pos.y);
        }
    }

    handleMouseUp(e) {
        if (this.isPanning) {
            this.isPanning = false;
            this.canvas.classList.remove('panning');
        } else if (this.isDrawing) {
            this.isDrawing = false;
            this.endDrawing();
        }
    }

    handleWheel(e) {
        e.preventDefault();

        // 줌 팩터 계산
        const zoomIntensity = 0.1;
        const delta = e.deltaY > 0 ? -zoomIntensity : zoomIntensity;
        const newScale = this.scale * (1 + delta);

        // 줌 제한
        if (newScale < 0.1 || newScale > 10) return;

        // 마우스 위치를 중심으로 줌
        this.zoom(newScale, e.clientX, e.clientY);

        // 줌 레벨 표시
        document.getElementById('zoom-level').textContent =
            `Zoom: ${Math.round(this.scale * 100)}%`;
    }

    handleTouchStart(e) {
        if (e.touches.length === 1) {
            const touch = e.touches[0];
            const pos = this.screenToCanvas(touch.clientX, touch.clientY);
            this.isDrawing = true;
            this.startDrawing(pos.x, pos.y);
        } else if (e.touches.length === 2) {
            this.isPanning = true;
            const touch = e.touches[0];
            this.lastX = touch.clientX;
            this.lastY = touch.clientY;
        }
    }

    handleTouchMove(e) {
        e.preventDefault();

        if (e.touches.length === 1 && this.isDrawing) {
            const touch = e.touches[0];
            const pos = this.screenToCanvas(touch.clientX, touch.clientY);
            this.draw(pos.x, pos.y);
        } else if (e.touches.length === 2 && this.isPanning) {
            const touch = e.touches[0];
            const dx = touch.clientX - this.lastX;
            const dy = touch.clientY - this.lastY;
            this.pan(dx, dy);
            this.lastX = touch.clientX;
            this.lastY = touch.clientY;
        }
    }

    handleTouchEnd(e) {
        if (this.isDrawing) {
            this.isDrawing = false;
            this.endDrawing();
        }
        if (this.isPanning) {
            this.isPanning = false;
        }
    }

    pan(dx, dy) {
        this.offsetX += dx;
        this.offsetY += dy;
        this.render();
    }

    zoom(newScale, centerX, centerY) {
        // 줌 중심점 계산
        const oldScale = this.scale;
        const scaleChange = newScale / oldScale;

        // 오프셋 조정 (중심점 기준)
        this.offsetX = centerX - (centerX - this.offsetX) * scaleChange;
        this.offsetY = centerY - (centerY - this.offsetY) * scaleChange;

        this.scale = newScale;
        this.render();
    }

    startDrawing(x, y) {
        this.currentStroke = {
            points: [{x, y}],
            color: this.penColor,
            size: this.penSize
        };
    }

    draw(x, y) {
        if (!this.currentStroke) return;

        this.currentStroke.points.push({x, y});
        this.render();
    }

    endDrawing() {
        if (this.currentStroke && this.currentStroke.points.length > 0) {
            this.strokes.push(this.currentStroke);
            this.currentStroke = null;
        }
    }

    render() {
        // 캔버스 클리어
        this.ctx.clearRect(0, 0, this.canvas.width, this.canvas.height);

        // 변환 적용
        this.ctx.save();
        this.ctx.translate(this.offsetX, this.offsetY);
        this.ctx.scale(this.scale, this.scale);

        // 배경 그리드 (선택사항)
        this.drawGrid();

        // 모든 스트로크 그리기
        this.strokes.forEach(stroke => this.drawStroke(stroke));

        // 현재 그리는 스트로크
        if (this.currentStroke) {
            this.drawStroke(this.currentStroke);
        }

        this.ctx.restore();
    }

    drawGrid() {
        const gridSize = 50;
        const startX = Math.floor(-this.offsetX / this.scale / gridSize) * gridSize;
        const startY = Math.floor(-this.offsetY / this.scale / gridSize) * gridSize;
        const endX = startX + this.canvas.width / this.scale + gridSize;
        const endY = startY + this.canvas.height / this.scale + gridSize;

        this.ctx.strokeStyle = '#f0f0f0';
        this.ctx.lineWidth = 1 / this.scale;

        for (let x = startX; x < endX; x += gridSize) {
            this.ctx.beginPath();
            this.ctx.moveTo(x, startY);
            this.ctx.lineTo(x, endY);
            this.ctx.stroke();
        }

        for (let y = startY; y < endY; y += gridSize) {
            this.ctx.beginPath();
            this.ctx.moveTo(startX, y);
            this.ctx.lineTo(endX, y);
            this.ctx.stroke();
        }
    }

    drawStroke(stroke) {
        if (stroke.points.length === 0) return;

        this.ctx.strokeStyle = stroke.color;
        this.ctx.lineWidth = stroke.size;
        this.ctx.lineCap = 'round';
        this.ctx.lineJoin = 'round';

        this.ctx.beginPath();
        const firstPoint = stroke.points[0];
        this.ctx.moveTo(firstPoint.x, firstPoint.y);

        for (let i = 1; i < stroke.points.length; i++) {
            const point = stroke.points[i];
            this.ctx.lineTo(point.x, point.y);
        }

        this.ctx.stroke();
    }

    clear() {
        this.strokes = [];
        this.currentStroke = null;
        this.render();
    }

    resetView() {
        this.offsetX = 0;
        this.offsetY = 0;
        this.scale = 1;
        this.render();
        document.getElementById('zoom-level').textContent = 'Zoom: 100%';
    }

    setPenSize(size) {
        this.penSize = size;
    }

    setPenColor(color) {
        this.penColor = color;
    }
}

// 앱 초기화
let whiteboard;

document.addEventListener('DOMContentLoaded', () => {
    const canvas = document.getElementById('whiteboard');
    whiteboard = new Whiteboard(canvas);

    // 펜 크기 버튼
    document.querySelectorAll('.pen-size').forEach(btn => {
        btn.addEventListener('click', () => {
            document.querySelectorAll('.pen-size').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            const size = parseInt(btn.dataset.size);
            whiteboard.setPenSize(size);
        });
    });

    // 색상 버튼
    document.querySelectorAll('.color-btn').forEach(btn => {
        btn.addEventListener('click', () => {
            document.querySelectorAll('.color-btn').forEach(b => b.classList.remove('active'));
            btn.classList.add('active');
            const color = btn.dataset.color;
            whiteboard.setPenColor(color);
        });
    });

    // 커스텀 색상
    document.getElementById('custom-color').addEventListener('input', (e) => {
        document.querySelectorAll('.color-btn').forEach(b => b.classList.remove('active'));
        whiteboard.setPenColor(e.target.value);
    });

    // 지우기 버튼
    document.getElementById('clear-btn').addEventListener('click', () => {
        if (confirm('정말로 모든 내용을 지우시겠습니까?')) {
            whiteboard.clear();
        }
    });

    // 리셋 버튼
    document.getElementById('reset-view-btn').addEventListener('click', () => {
        whiteboard.resetView();
    });

    // 키보드 단축키
    document.addEventListener('keydown', (e) => {
        // Space: 팬 모드 표시
        if (e.code === 'Space' && !e.repeat) {
            canvas.classList.add('panning');
        }

        // Ctrl/Cmd + Z: 실행 취소 (향후 구현)
        if ((e.ctrlKey || e.metaKey) && e.key === 'z') {
            e.preventDefault();
            // TODO: 실행 취소 기능
        }
    });

    document.addEventListener('keyup', (e) => {
        if (e.code === 'Space') {
            canvas.classList.remove('panning');
        }
    });
});

// 테스트를 위한 export (Node.js 환경)
if (typeof module !== 'undefined' && module.exports) {
    module.exports = Whiteboard;
}
