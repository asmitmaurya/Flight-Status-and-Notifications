from flask import Flask, jsonify, request
from pymongo import MongoClient
from datetime import datetime
import firebase_admin
from firebase_admin import credentials, messaging

# Initialize Flask app
app = Flask(__name__)

# Initialize Firebase Admin SDK (replace with your own service account key)
cred = credentials.Certificate('path/to/serviceAccountKey.json')
firebase_admin.initialize_app(cred)

# Initialize MongoDB client
client = MongoClient('mongodb://localhost:27017/')
db = client['flight_status']
flights_collection = db['flights']

# Mock data for flights (replace with actual data fetching from airport system)
flights = [
    {"flight_number": "ABC123", "status": "On time", "gate": "A1"},
    {"flight_number": "XYZ789", "status": "Delayed", "gate": "B2"}
]

# Endpoint to get all flights
@app.route('/api/flights', methods=['GET'])
def get_flights():
    return jsonify({"flights": flights})

# Endpoint to update flight status (mocking real-time update)
@app.route('/api/update_status/<flight_number>', methods=['PUT'])
def update_status(flight_number):
    new_status = request.json.get('status')
    for flight in flights:
        if flight['flight_number'] == flight_number:
            flight['status'] = new_status
            # Send notification when status changes
            send_notification(flight_number, new_status)
            break
    return jsonify({"message": "Flight status updated."})

# Function to send notification using Firebase Cloud Messaging (FCM)
def send_notification(flight_number, new_status):
    registration_token = 'YOUR_DEVICE_TOKEN_HERE'
    message = messaging.Message(
        data={
            'flight_number': flight_number,
            'status': new_status,
            'timestamp': str(datetime.now())
        },
        token=registration_token,
    )
    response = messaging.send(message)
    print('Notification sent:', response)

if __name__ == '__main__':
    app.run(debug=True)
