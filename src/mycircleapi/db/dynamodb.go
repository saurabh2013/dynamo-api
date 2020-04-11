package db

import (
	"fmt"
	model "mycircleapi/models"
	"strconv"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

const TblRegistration = "Registration"
const TblContactList = "ContactList"
const TblAffectedList = "AffectedList"

// IdAttribute - attribute name for the partition key.
const IdAttribute = "deviceid"
const MobileAttribute = "mobile"

type Dynamodb struct {
	*dynamodb.DynamoDB
}

var Items Dynamodb

func InitializeDB(region, endpoint string) error {
	fmt.Println("Initializing DynamoDB..")
	// Initialize the AWS session.
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String(region),
		Endpoint: aws.String(endpoint),
	})
	if err != nil {
		return fmt.Errorf("INITIALIZATION ERROR: %v", err)
	}

	// Initialize the DynamoDB instance.
	Items = Dynamodb{dynamodb.New(sess, aws.NewConfig().WithLogLevel(aws.LogDebugWithHTTPBody))}
	//Items.listTables()

	tables := []string{TblRegistration, TblContactList, TblAffectedList}
	for _, t := range tables {
		fmt.Println("Testinng Table", t)
		tableExists, err := Items.tableExists(t)
		if err != nil {
			return fmt.Errorf("INITIALIZATION ERROR: %v", err)
		}

		if !tableExists {
			if t == TblAffectedList {
				createTable(t, MobileAttribute)

			} else {
				createTable(t, IdAttribute)
			}
		} else {
			fmt.Println("Table already exists!")
		}
	}
	return nil
}

// Contact - adds a new contact to the database.
func (db Dynamodb) AddContact(newcontact model.Contact) error {
	fmt.Printf("++%v", newcontact)
	data, err := dynamodbattribute.MarshalMap(newcontact)
	if err != nil {
		return fmt.Errorf("AddProduct -> Error marshalling product: %v", err)
	}

	// Setup the insert criteria.
	item := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(TblContactList),
	}

	// Insert the new Product into the database.
	_, err = Items.PutItem(item)
	if err != nil {
		return fmt.Errorf("AddProduct -> New product could not be added: %v", err)
	}

	return nil
}

// Contact - adds a new contact to the database.
func (db Dynamodb) AddUpdateAffected(req model.Affected) error {
	fmt.Printf("++%v", req)
	// data, err := dynamodbattribute.MarshalMap(newcontact)
	// if err != nil {
	// 	return fmt.Errorf("AddProduct -> Error marshalling product: %v", err)
	// }

	// // Setup the insert criteria.
	// item := &dynamodb.PutItemInput{
	// 	Item:      data,
	// 	TableName: aws.String(TblContactList),
	// }

	// // Insert the new Product into the database.
	// _, err = Items.PutItem(item)
	// if err != nil {
	// 	return fmt.Errorf("AddProduct -> New product could not be added: %v", err)
	// }

	///
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(TblAffectedList),
		Key: map[string]*dynamodb.AttributeValue{
			MobileAttribute: {S: aws.String(req.Mobile)},
		},
		UpdateExpression:         aws.String("SET mobile = :mobile, affectedstatus = :affectedstatus"),
		ExpressionAttributeNames: map[string]*string{"mobile": aws.String("mobile")},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":mobile":         {S: aws.String(req.Mobile)},
			":affectedstatus": {N: aws.String(strconv.Itoa(req.AffectedStatus))},
		},
		ReturnValues: aws.String("ALL_NEW"),
	}

	// Execute the update.
	_, err := Items.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("New item <%v> could not be updated/added: %v", req, err)
	}

	return nil
}

// Add device - adds a new device to the database.
func (db Dynamodb) AddDevice(newcontact model.Registration) error {
	fmt.Printf("++%v", newcontact)
	data, err := dynamodbattribute.MarshalMap(newcontact)
	if err != nil {
		return fmt.Errorf("AddDevice -> Error marshalling device: %v", err)
	}

	// Setup the insert criteria.
	item := &dynamodb.PutItemInput{
		Item:      data,
		TableName: aws.String(TblRegistration),
	}

	// Insert the new Product into the database.
	_, err = Items.PutItem(item)
	if err != nil {
		return fmt.Errorf("AddDevice -> New device could not be added: %v", err)
	}

	return nil
}

// GetContact - if it exists, retrieves the requested Product from the database;
func (db Dynamodb) GetContact(id string) (c model.Contact, e error) {
	// Setup query criteria.
	result, err := Items.Query(&dynamodb.QueryInput{
		TableName:              aws.String(TblContactList),
		ScanIndexForward:       aws.Bool(false),
		KeyConditionExpression: aws.String("deviceid = :deviceid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":deviceid": {S: aws.String(id)},
		},
	})

	// If the contact was found, then there should only be one item.
	for _, i := range result.Items {
		var p model.Contact
		err = dynamodbattribute.UnmarshalMap(i, &p)
		if err != nil {
			return c, fmt.Errorf("Unmarshalling Contact failed:\n%v", err)
		}

		c = p

		return
	}

	return c, fmt.Errorf("Contact <%v> does not exist", id)
}

// GetRegistration - if it exists, retrieves the requested Product from the database;
func (db Dynamodb) GetRegistration(id string) (c model.Registration, e error) {
	// Setup query criteria.
	result, err := Items.Query(&dynamodb.QueryInput{
		TableName:              aws.String(TblRegistration),
		ScanIndexForward:       aws.Bool(false),
		KeyConditionExpression: aws.String("deviceid = :deviceid"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":deviceid": {S: aws.String(id)},
		},
	})

	// If the contact was found, then there should only be one item.
	for _, i := range result.Items {
		var p model.Registration
		err = dynamodbattribute.UnmarshalMap(i, &p)
		if err != nil {
			return c, fmt.Errorf("Unmarshalling Registration failed:\n%v", err)
		}

		c = p

		return
	}

	return c, fmt.Errorf("Registration <%v> does not exist", id)
}

// GeAffectedn -
func (db Dynamodb) GeAffected(id string) (c model.Affected, e error) {
	// Setup query criteria.
	result, err := Items.Query(&dynamodb.QueryInput{
		TableName:              aws.String(TblAffectedList),
		ScanIndexForward:       aws.Bool(false),
		KeyConditionExpression: aws.String("mobile = :mobile"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":mobile": {S: aws.String(id)},
		},
	})

	// If the contact was found, then there should only be one item.
	for _, i := range result.Items {
		var p model.Affected
		err = dynamodbattribute.UnmarshalMap(i, &p)
		if err != nil {
			return c, fmt.Errorf("Unmarshalling Affectedlist failed:\n%v", err)
		}

		c = p

		return
	}

	return c, fmt.Errorf("Affectedlist <%v> does not exist", id)
}

// DeleteContact - if it exists, deletes the specified Product.
func (db Dynamodb) Delete(id, idAttribute, table string) error {
	// Setup the delete criteria.
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(table),
		Key: map[string]*dynamodb.AttributeValue{
			idAttribute: {S: aws.String(id)},
		},
		ReturnValues: aws.String("ALL_OLD"),
	}

	// Process the deletion.
	results, err := Items.DeleteItem(input)
	if err != nil {
		return fmt.Errorf("Item <%v> could not be deleted: %v", id, err)
	}

	// If there was nothing to delete, then return an appropriate message.
	if len(results.Attributes) == 0 {
		return fmt.Errorf("Item <%v> does not exist", id)
	}

	return nil
}

func Cleanup() error {
	fmt.Println("Cleaning up...")
	// TODO - Delete table and dynamic resources
	return nil
}

func createTable(t, id string) error {
	fmt.Println("Creating table...")

	// Setup table create criteria.
	input := &dynamodb.CreateTableInput{
		TableName: aws.String(t),
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String(id), KeyType: aws.String("HASH"),
			},
		},
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String(id), AttributeType: aws.String("S"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits: aws.Int64(10), WriteCapacityUnits: aws.Int64(10),
		},
	}

	// Create the table.
	if _, err := Items.CreateTable(input); err != nil {
		fmt.Println("Error during CreateTable:")
		return fmt.Errorf("%v", err)
	}

	fmt.Printf("Table '%v' successfully created!\n", TblContactList)

	// Initialize the database with some data for testing purposes.
	//enterTestData()

	return nil
}

// enterTestData - local helper function that populates the database with some dummy data for testing purposes.
func enterTestData() error {
	contacts := []model.Contact{{DeviceId: "testdevice", Name: "samtest"}}

	for _, p := range contacts {
		err := Items.AddContact(p)
		if err != nil {
			return fmt.Errorf("Error entering test data: %v", err)
		}
	}

	return nil
}

//delete table for testing
func (db Dynamodb) deletecontactTable(name string) error {

	t := &dynamodb.DeleteTableInput{TableName: aws.String(name)}

	_, e := db.DeleteTable(t)
	return e
}

func (db Dynamodb) tableExists(name string) (bool, error) {
	result, err := db.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println("Error during ListTables:")
		return false, fmt.Errorf("%v", err)
	}

	for _, n := range result.TableNames {
		if *n == name {

			return true, nil
		}
	}

	return false, nil
}

func (db Dynamodb) listTables() error {
	fmt.Print("Getting list of tables")
	result, err := db.ListTables(&dynamodb.ListTablesInput{})

	if err != nil {
		fmt.Println("Error during ListTables:")
		return fmt.Errorf("%v", err)
	}

	fmt.Println("Tables:")
	fmt.Println("")

	for _, n := range result.TableNames {
		fmt.Println(*n)
	}
	return nil
}

func (db Dynamodb) UpdateEffected(ph string, value int) error {
	// Setup the update criteria.
	input := &dynamodb.UpdateItemInput{
		TableName: aws.String(TblContactList),
		Key: map[string]*dynamodb.AttributeValue{
			IdAttribute: {S: aws.String(ph)},
		},
		UpdateExpression: aws.String("SET healthstatus = :healthstatus"),
		//ExpressionAttributeNames: map[string]*string{"healthstatus": aws.String("healthstatus")},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":healthstatus": {N: aws.String(strconv.Itoa(value))},
		},
		ReturnValues: aws.String("UPDATED_NEW"),
	}

	// Execute the update.
	_, err := Items.UpdateItem(input)
	if err != nil {
		return fmt.Errorf("Contact for <%v> could not be updated: %v", ph, err)
	}

	return nil
}
