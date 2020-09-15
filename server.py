from flask import Flask
import sys
app = Flask(__name__)

@app.route('/')
def hello():
    return "Server: "+sys.argv[1]

if __name__ == '__main__':
    app.run(port=sys.argv[1])