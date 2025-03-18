To  run the program use below command

go run main.go


Tested Request Body
1. First
schema
  @link(url: "https://specs.apollo.dev/link/v1.0")
  @link(url: "https://specs.apollo.dev/join/v0.3", for: EXECUTION)
{
  query: Query
}

directive @join__enumValue(graph: join__Graph!) repeatable on ENUM_VALUE

directive @join__field(graph: join__Graph, requires: join__FieldSet, provides: join__FieldSet, type: String, external: Boolean, override: String, usedOverridden: Boolean) repeatable on FIELD_DEFINITION | INPUT_FIELD_DEFINITION

directive @join__graph(name: String!, url: String!) on ENUM_VALUE

directive @join__implements(graph: join__Graph!, interface: String!) repeatable on OBJECT | INTERFACE

directive @join__type(graph: join__Graph!, key: join__FieldSet, extension: Boolean! = false, resolvable: Boolean! = true, isInterfaceObject: Boolean! = false) repeatable on OBJECT | INTERFACE | UNION | ENUM | INPUT_OBJECT | SCALAR

directive @join__unionMember(graph: join__Graph!, member: String!) repeatable on UNION

directive @link(url: String, as: String, for: link__Purpose, import: [link__Import]) repeatable on SCHEMA

type Account
  @join__type(graph: SUBGRAPH)
{
  accountReferenceId: String
  availableCreditAmount: AvailableCreditAmount
}

type AvailableCreditAmount
  @join__type(graph: SUBGRAPH)
{
  availableSpendingCreditAmount: Float
  availableCashCreditAmount: Float
}

type Customer
  @join__type(graph: SUBGRAPH)
{
  name: String
  age: Int
  accounts: [Account]
}

type DataProduct
  @join__type(graph: SUBGRAPH)
{
  dataProductIdentifier: String
  attribute3: EntityAttribute
}

type EntityAttribute
  @join__type(graph: SUBGRAPH)
{
  entityAttribute1: String
  entityAttribute2: String
}

scalar join__FieldSet

enum join__Graph {
  SUBGRAPH @join__graph(name: "subgraph", url: "http://localhost:4001")
}

scalar link__Import

enum link__Purpose {
  """
  `SECURITY` features provide metadata necessary to securely resolve fields.
  """
  SECURITY

  """
  `EXECUTION` features provide metadata necessary for operation execution.
  """
  EXECUTION
}

type Query
  @join__type(graph: SUBGRAPH)
{
  accounts: [Account]
  customerById(customerReferenceId: String!): Customer
  dataProducts: [DataProduct]
  dataProduct(dataProductIdentifier: String!): DataProduct
}


2. Second
# Define the types for the schema

# This type represents the attributes of a data product with entity attributes.
type EntityAttribute {
  entityAttribute1: String
  entityAttribute2: String
}

# This type represents a data product with a unique identifier and associated attributes.
type DataProduct {
  dataProductIdentifier: String
  attribute3: EntityAttribute
}

# Define the root Query type
type Query {
  # A query to fetch all data products
  dataProducts: [DataProduct]

  # A query to fetch a specific data product by its identifier
  dataProduct(dataProductIdentifier: String!): DataProduct
}

3. Third
type AvailableCreditAmount {
  availableSpendingCreditAmount: Float
  availableCashCreditAmount: Float
}

type Account {
  accountReferenceId: String
  availableCreditAmount: AvailableCreditAmount
}

type Customer {
  name: String
  age: Int
  accounts: [Account]
}

type Query {
  customerById(customerReferenceId: String!): Customer
}



