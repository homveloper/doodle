---
trigger: always_on
description: 코드 생성, 리팩토링, 함수/컴포넌트/훅 작성 시 적용. 신규 파일 생성, 복잡한 로직 구현, 비즈니스 규칙 코딩, API 연동, 상태관리 코드 작성 시 트리거. 코드를 처음 보는 개발자가 도메인 지식까지 습득할 수 있도록 친절하고 교육적인 주석을 상세히 작성해야 함. 주석이 곧 문서이자 튜토리얼 역할.
globs: *.js, *.ts, *.tsx
---

{
  "rule_name": "Talkative Mode: 수다쟁이 교육자 스타일 주석 규칙",
  "version": "1.0.0",
  "description": "코드 생성, 리팩토링, 함수/컴포넌트/훅 작성 시 적용. 신규 파일 생성, 복잡한 로직 구현, 비즈니스 규칙 코딩, API 연동, 상태관리 코드 작성 시 트리거. 코드를 처음 보는 개발자가 도메인 지식까지 습득할 수 있도록 친절하고 교육적인 주석을 상세히 작성. 주석이 곧 문서이자 튜토리얼 역할.",
  "philosophy": {
    "mindset": "주석은 미래의 동료에게 보내는 친절한 편지다. 6개월 후의 나 자신도 '처음 보는 사람'이라고 생각하라.",
    "roles": [
      {
        "role": "교육자",
        "behavior": "이걸 왜 이렇게 했을까?라는 질문에 미리 답한다"
      },
      {
        "role": "수다쟁이",
        "behavior": "혼잣말하듯 생각의 흐름을 자연스럽게 풀어쓴다"
      },
      {
        "role": "스토리텔러",
        "behavior": "코드가 탄생한 배경과 맥락을 이야기한다"
      },
      {
        "role": "멘토",
        "behavior": "도메인 지식을 자연스럽게 전달한다"
      }
    ]
  },
  "core_principles": [
    {
      "principle": "맥락 제공",
      "description": "해당 코드베이스를 처음 보는 개발자도 쉽게 이해할 수 있도록 충분한 맥락을 제공한다."
    },
    {
      "principle": "Why 중심",
      "description": "코드가 '무엇을 하는지(What)'뿐만 아니라 '왜 이렇게 하는지(Why)'를 반드시 설명한다."
    },
    {
      "principle": "자기문서화",
      "description": "주석을 통해 코드 자체가 문서이자 튜토리얼 역할을 하도록 한다."
    },
    {
      "principle": "도메인 교육",
      "description": "비즈니스 용어, 업계 관행, 기술적 배경지식을 주석 내에서 자연스럽게 설명한다."
    }
  ],
  "comment_guide": {
    "file_module_level": {
      "scope": "파일 상단, 모듈 전체",
      "goal": "파일을 열었을 때 '아, 이 파일은 이런 역할이구나!' 하고 바로 감이 오도록 한다",
      "required_sections": [
        {
          "section": "한 줄 요약",
          "description": "이 모듈이 뭔지 한 문장으로"
        },
        {
          "section": "이 파일은 뭘 하나요?",
          "description": "2-3문장으로 존재 이유와 핵심 책임 설명"
        },
        {
          "section": "왜 이렇게 분리했나요?",
          "description": "아키텍처적 결정 배경, 다른 모듈과의 관계"
        },
        {
          "section": "알아두면 좋은 도메인 지식",
          "description": "이 코드를 이해하는 데 필요한 비즈니스/기술 배경을 친절하게 풀어서 설명"
        },
        {
          "section": "주요 의존성",
          "description": "의존하는 모듈과 왜 의존하는지"
        },
        {
          "section": "함께 보면 좋은 파일",
          "description": "관련 파일과 어떤 맥락에서 연결되는지"
        }
      ],
      "tone_guide": "마치 후임 개발자에게 온보딩하듯이, '이 파일 처음 보죠? 제가 설명해드릴게요~' 느낌으로 작성"
    },
    "function_method_level": {
      "scope": "함수, 메서드, 커스텀 훅",
      "goal": "함수를 보고 '이 함수 쓰면 되겠다!' 또는 '이건 내 상황에 안 맞겠네'를 바로 판단할 수 있도록 한다",
      "required_sections": [
        {
          "section": "한 줄 요약",
          "description": "동사로 시작하는 한 문장 (예: '주문 금액에 적용 가능한 최적의 할인을 계산합니다')"
        },
        {
          "section": "언제 이 함수를 쓰나요?",
          "description": "사용 시나리오, 어떤 상황에서 호출하는지 구체적으로"
        },
        {
          "section": "어떻게 동작하나요?",
          "description": "내부 로직을 단계별로 설명. 복잡하다면 1️⃣ 2️⃣ 3️⃣ 번호 매기기"
        },
        {
          "section": "왜 이렇게 구현했나요?",
          "description": "설계 결정의 이유를 Q&A 형식으로. 대안이 있었다면 왜 이걸 선택했는지"
        },
        {
          "section": "@param",
          "description": "파라미터 설명 + 예시값 필수"
        },
        {
          "section": "@returns",
          "description": "반환값 설명 + 어떤 형태인지, 어떻게 활용하면 좋은지"
        },
        {
          "section": "@throws",
          "description": "어떤 상황에서 에러가 나는지"
        },
        {
          "section": "@example",
          "description": "실제 사용 예시 코드 (복사해서 바로 쓸 수 있게)"
        }
      ],
      "tone_guide": "마치 페어 프로그래밍하면서 '이 함수는요~' 하고 설명해주는 느낌으로"
    },
    "logic_level": {
      "scope": "인라인 주석, 조건문, 반복문, 복잡한 표현식",
      "goal": "코드를 읽다가 '응? 이게 왜 여기 있지?' 싶은 부분에 친절한 설명을 단다",
      "required_elements": [
        {
          "element": "비즈니스 로직 의도",
          "description": "이 조건이 왜 필요한지, 어떤 비즈니스 규칙을 반영하는지"
        },
        {
          "element": "분기 이유",
          "description": "if/else, switch 등에서 각 분기가 어떤 케이스를 처리하는지"
        },
        {
          "element": "매직 넘버 설명",
          "description": "숫자나 문자열 상수가 왜 그 값인지 (예: '30일 = 한 달 기준, 기획팀 정책')"
        },
        {
          "element": "함정 경고",
          "description": "주의해야 할 점, 과거에 실수했던 부분, 흔히 하는 오해"
        },
        {
          "element": "역사적 맥락",
          "description": "왜 이런 코드가 생겼는지 히스토리 (예: '2023년 8월 봇 공격 이후 추가됨')"
        }
      ],
      "formatting": {
        "section_separator": "═══════════════════════════════════════════════════════════",
        "step_indicator": "🎯 STEP N: [단계명]",
        "warning_prefix": "⚠️ 주의:",
        "important_prefix": "🔥 중요:",
        "tip_prefix": "💡 팁:",
        "question_prefix": "Q:",
        "answer_prefix": "A:"
      },
      "tone_guide": "마치 코드 리뷰하면서 '여기 이렇게 한 이유가요~' 하고 설명해주는 느낌으로"
    }
  },
  "domain_knowledge_injection": {
    "description": "코드를 처음 보는 사람이 도메인 지식을 자연스럽게 습득할 수 있도록 한다",
    "techniques": [
      {
        "technique": "용어 정의 삽입",
        "example": "PG사(Payment Gateway)란? 우리 서비스와 실제 카드사/은행 사이를 연결해주는 중개 업체입니다."
      },
      {
        "technique": "비유 활용",
        "example": "Redis의 pub/sub은 카카오톡 단체방 같은 거예요. 메시지 보내면 방에 있는 모든 사람이 받죠."
      },
      {
        "technique": "실제 사례 언급",
        "example": "이 검증 로직은 2023년 8월 봇 공격 사건 이후 추가됐어요. 그때 가짜 주문이 10만 건..."
      },
      {
        "technique": "대안 비교",
        "example": "REST 대신 GraphQL을 쓴 이유는, 모바일에서 필요한 필드만 골라 받아 데이터 사용량을 줄이려고요."
      },
      {
        "technique": "흔한 오해 해소",
        "example": "JWT가 암호화라고 오해하기 쉬운데, 사실 서명일 뿐이에요. 내용은 누구나 볼 수 있습니다!"
      }
    ]
  },
  "constraints": [
    {
      "rule": "명백한 코드에 불필요한 주석 금지",
      "bad_example": "i++ // i를 1 증가",
      "good_example": "i++ // 다음 페이지로 이동 (페이지네이션 인덱스)"
    },
    {
      "rule": "주석과 코드 동기화 필수",
      "description": "코드 수정 시 관련 주석도 반드시 함께 업데이트한다"
    },
    {
      "rule": "거짓 정보 금지",
      "description": "확실하지 않은 내용은 '추정', '아마도' 등을 명시한다"
    },
    {
      "rule": "개인 비하 금지",
      "bad_example": "// 이전 개발자가 이상하게 짜놔서...",
      "good_example": "// 레거시 호환성을 위해 이 방식을 유지합니다"
    }
  ]
}