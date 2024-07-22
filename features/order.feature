Feature: Order
  Scenario: I can place an order
    When make a POST request to "/orders"
    """
    {
      "name": "Aotoki",
      "items": [
        {
          "name": "Apple",
          "quantity": 2,
          "unit_price": 10
        },
        {
          "name": "Banana",
          "quantity": 3,
          "unit_price": 5
        }
      ]
    }
    """
    Then the response status code should be 200
    And the response JSON contains "id" string
    And the response JSON contains "name" with value "Aotoki"
    And the response JSON contains "items[0].name" with value "Apple"
    And the response JSON contains "items[0].quantity" with value 2
    And the response JSON contains "items[0].unit_price" with value 10
    And the response JSON contains "items[1].name" with value "Banana"
    And the response JSON contains "items[1].quantity" with value 3
    And the response JSON contains "items[1].unit_price" with value 5

  Scenario: Cannot place order with same item name
    When make a POST request to "/orders"
    """
    {
      "name": "Aotoki",
      "items": [
        {
          "name": "Apple",
          "quantity": 2,
          "unit_price": 10
        },
        {
          "name": "Apple",
          "quantity": 3,
          "unit_price": 5
        }
      ]
    }
    """
    Then the response status code should be 400
    And the response body should be "item name must be unique"
