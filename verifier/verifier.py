#!flask/bin/python
from flask import Flask
from flask import request
import jwt

app = Flask(__name__)

@app.route('/verifier', methods=['POST'])
def verifier():
    if "token" in request.json:
        token = request.json['token']
        print(token)
        try:
            file = open("/usr/share/key/key.txt", "r")
            output = jwt.decode(token, file.read())
            return output['username']
        except Exception as e:
            return str(e)
    else:
        return "bad"

if __name__ == '__main__':
    app.run(host='0.0.0.0')