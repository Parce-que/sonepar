Scenario Outline: Looking for words

Given I am on a new pages of "searchEngine"
When I look for "words"
And I click the search button
Then I should see the "words" in the resulting page

Examples :
            | words           | searchEngine |
            | oiseau          | google       |
            | Serendipity     | google       |
            | Stephen Hawking | google       |
            | golang          | google       |