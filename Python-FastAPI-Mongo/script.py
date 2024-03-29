import pymongo
import json

# Connect to MongoDB
client = pymongo.MongoClient("mongodb://localhost:27017")
db = client["courses"]
collection = db["courses"]

# Read courses from courses.json
with open("courses.json","r") as f:
    courses = json.load(f)

# Create index for efficient retrival
collection.create_index("name")

# Add rating field to each course
for course in courses:
    course['rating'] = {'total':0,'count':0}

# Add rating field to each chapter
for course in courses:
    for chapter in course['chapters']:
        chapter['rating'] = {'total':0,'count':0}

# Add courses to collection
for course in courses:
    collection.insert_one(course)

# Fetching Data From Database
coursesData = db.get_collection("courses")
for course in coursesData.find():
    courseRating = course["rating"][0]
    print(courseRating)

#Close MongoDB connection
client.close()