from datetime import datetime

from flask import Flask, request, jsonify

app = Flask(__name__)


@app.route("/process_question", methods=["POST"])
def process_question():
    data = request.get_json()
    question = data.get("question")

    # Proses pertanyaan dengan model AI
    response = ai_model_function(question)

    return jsonify({"answer": response})


def ai_model_function(question):
    if "cuaca" in question:
        return "Hari ini cerah."
    elif "waktu" or "jam" in question:
        now = datetime.now()
        current_time = now.strftime("%H:%M:%S")
        return f"Waktunya sekarang jam {current_time} UTC+7"
    else:
        return "Maaf, saya tidak mengerti pertanyaan Anda."


if __name__ == "__main__":
    app.run(debug=True)
