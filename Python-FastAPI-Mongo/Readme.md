### Installation

1. Installing dependencies `pip install fastapi pymongo uvicorn starlette pydantic`
2. To install packages individually
```
pip install pymongo
pip 
```
3. Execute `python3 script.py`
3. `uvicorn main:app --reload`
4. Go to browser hit: <br/>
   [] `http://localhost:8000/` <br/>
   [] `http://127.0.0.1:8000/courses?sort=desc`
5. To view swagger document.
   `http://127.0.0.1:8000/docs`
6. To view better docs
   `http://127.0.0.1:8000/redoc`
7. To get specific course
   `http://127.0.0.1:8000/courses/66078d3122ae5cb883b0c46c`
   To get above id, need to connect with mongo client.