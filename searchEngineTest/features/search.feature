Scenario Outline: Looking for words

Given I am on a new pages of "searchEngine"
When I look for "words"
Then I should have a "200" http code
And I should see the "words" in the resulting page

Examples :
            | words           | searchEngine |
            | oiseau          | google       |
            | Serendipity     | google       |
            | Stephen Hawking | google       |
            | golang          | google       |