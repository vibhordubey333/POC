import contextlib
from fastapi import FastAPI, HTTPException, Query
from pymongo import MongoClient
from bson import ObjectId
from fastapi.encoders import jsonable_encoder

app = FastAPI()
client = MongoClient('mongodb://localhost:27017')
db = client['courses']


@app.get("/")
def index():
    return {"Welcome", "Tourist"}


@app.get('/courses')
def get_courses(sort_by: str = 'date', domain: str = None):
    # set the rating.total and rating.count to all the courses based on the sum of the chapters rating
    for course in db.courses.find():
        total = 0
        count = 0
        for chapter in course['chapters']:
            with contextlib.suppress(KeyError):
                total += chapter['rating']['total']
                count += chapter['rating']['count']
        db.courses.update_one({'_id': course['_id']}, {'$set': {'rating': {'total': total, 'count': count}}})


    # sort_by == 'date' [DESCENDING]
    if sort_by == 'date':
        sort_field = 'date'
        sort_order = -1

    # sort_by == 'rating' [DESCENDING]
    elif sort_by == 'rating':
        sort_field = 'rating.total'
        sort_order = -1

    # sort_by == 'alphabetical' [ASCENDING]
    else:
        sort_field = 'name'
        sort_order = 1

    query = {}
    if domain:
        query['domain'] = domain


    courses = db.courses.find(query, {'name': 1, 'date': 1, 'description': 1, 'domain':1,'rating':1,'_id': 0}).sort(sort_field, sort_order)
    return list(courses)

# @app.get('/courses/{course_id}')
# def get_course(course_id: str):
#     # Retrieve course data from the database based on the provided course_id
#     course = db.courses.find({'id':ObjectId(course_id)},{'_id':0,'chapters':0})
#     # If the course is not found in the db, raise am HTTPException with status code 404
#     if not course:
#         raise HTTPException(status_code = 404,detail='COURSE_NOT_FOUND')
#     # Try to extract the `total` rating from the 'rating' field of the course data
#     try:
#         course['rating']= course['rating']['total']
#     # If total key doesn't exists in the rating field, set 'rating' to 'Not Rated'
#     except KeyError:
#         course['rating'] = 'Not rated yet'
#     return course

@app.get('/courses/{course_id}')
def get_course(course_id: str):
    course = db.courses.find_one({'_id': ObjectId(course_id)}, {'_id': 0, 'chapters': 0})
    if not course:
        raise HTTPException(status_code=404, detail='Course not found')
    try:
        course['rating'] = course['rating']['total']
    except KeyError:
        course['rating'] = 'Not rated yet'
    return course