import os
from dotenv import load_dotenv
from openai import OpenAI

# 환경 변수 로드
load_dotenv()

# OpenAI API 키
api_key = os.environ.get("OPENAI_API_KEY")

# OpenAI 클라이언트 초기화
client = OpenAI(
    base_url="https://openrouter.ai/api/v1",
    api_key=api_key,
    )

def get_chat_completion(prompt, model="google/gemma-3-27b-it:free"):
    # OpenAI ChatCompletion API 호출
    response = client.chat.completions.create(
        model=model,
        messages=[
            {"role": "system", "content": "You are a helpful assistant."},
            {"role": "user", "content": prompt}
        ]
    )
    return response.choices[0].message.content

if __name__ == "__main__":

    # 사용자 입력
    user_input = input("User: ")

    # 응답 출력
    response = get_chat_completion(user_input)
    print("Assistant:", response)