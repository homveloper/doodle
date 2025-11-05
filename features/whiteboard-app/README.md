# 🎨 Whiteboard App - 무한 화이트보드

JavaScript로 구현한 무한 캔버스 그림판 애플리케이션

## 기능

- ✨ **무한 캔버스**: 팬(드래그)과 줌(스크롤) 지원
- 🖊️ **그리기 도구**: 다양한 펜 크기
- 🎨 **색상 선택**: 커스텀 색상 팔레트
- 💫 **플로팅 UI**: 하단에 떠있는 도구 패널
- 🎯 **부드러운 그리기**: Canvas API 활용

## 사용 방법

### 실행

```bash
# 브라우저에서 열기
open index.html
# 또는
python3 -m http.server 8000
# 그 다음 http://localhost:8000 접속
```

### 조작법

- **그리기**: 마우스 클릭 + 드래그
- **이동**: Space + 마우스 드래그 (또는 오른쪽 버튼)
- **줌**: 마우스 휠
- **펜 크기**: 하단 패널에서 선택
- **색상**: 하단 패널에서 선택

## 기술 스택

- Pure JavaScript (ES6+)
- HTML5 Canvas API
- CSS3 (Flexbox, Animations)
- No dependencies!

## 구조

```
whiteboard-app/
├── index.html      # 메인 HTML
├── style.css       # 스타일시트
├── whiteboard.js   # 메인 로직
├── test.html       # 간단한 테스트
└── README.md       # 이 파일
```

## 주요 클래스

### Whiteboard
- `pan(dx, dy)`: 캔버스 이동
- `zoom(scale, x, y)`: 줌 인/아웃
- `startDrawing(x, y)`: 그리기 시작
- `draw(x, y)`: 그리기
- `endDrawing()`: 그리기 종료

## 개발

```bash
# 테스트 실행
open test.html
```

## 향후 계획

- [ ] 실행 취소/다시 실행
- [ ] 도형 그리기 (사각형, 원, 선)
- [ ] 지우개 도구
- [ ] 저장/불러오기
- [ ] 여러 레이어
- [ ] 텍스트 입력

## 라이선스

MIT
