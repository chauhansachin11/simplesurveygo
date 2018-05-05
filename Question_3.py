from jsonschema import validate

schema = {
		  "colors": [
		    {
		        "color": {"type" : "string"},
		        "category": {"type" : "string"},
		        "type": {"type" : "string"},
		        "code": {
		        "rgba": [
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"}
                    ],
		        "hex": {"type" : "string","pattern": "^#[0-9][A-Z]$" }
		      }
		    },
		    {
		        "color": {"type" : "string"},
		        "category": {"type" : "string"},
		        "code": {
		        "rgba": [
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"}
                    ],
		        "hex": {"type" : "string","pattern": "^#[0-9][A-Z]$" }
		      }
		    },
		    {
		        "color": {"type" : "string"},
		        "category": {"type" : "string"},
		        "type": {"type" : "string"},
		        "code": {
		        "rgba": [
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"},
                    {"type" : "integer"}
                    ],
		        "hex": {"type" : "string","pattern": "^#[0-9][A-Z]$" }
		      }
		    },
		  ]
} 

def validate_schema(payload):
    try:
        validate(payload,schema)
        return True
    except:
        return False
