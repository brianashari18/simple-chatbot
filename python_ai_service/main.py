from flask import Flask, request, jsonify

app = Flask(__name__)


# Fungsi sederhana untuk menangani permintaan AI
@app.route("/process_question", methods=["POST"])
def process_question():
    data = request.get_json()
    question = data.get("question")

    # Proses pertanyaan dengan model AI
    response = ai_model_function(question)

    return jsonify({"answer": response})


def ai_model_function(question):
    # Logika AI di sini
    if "cuaca" in question:
        return "Hari ini cerah."
    elif "waktu" in question:
        return "Waktunya sekarang jam 14:00."
    else:
        return "Maaf, saya tidak mengerti pertanyaan Anda."


if __name__ == "__main__":
    app.run(debug=True)
