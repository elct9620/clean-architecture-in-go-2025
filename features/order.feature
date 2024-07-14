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
