Feature: Search Engine
    In order to look for thing one the web
    As a customer
    I need to be able to input words in a search engine

Scenario Outline: Looking for words
    Given I am on a new pages of "searchEngine"
    When I look for "words"
    And I click the search button
    Then I should see the "words" in the resulting page

Examples:
            | words           | searchEngine |
            | oiseau          | google       |
            | Serendipity     | google       |
            | Stephen Hawking | google       |
            | golang          | google       |