#include "VectorTestCase.hpp"

VectorTestCase::VectorTestCase() {}

VectorTestCase::VectorTestCase( const VectorTestCase & src ){ (void)src; }

VectorTestCase::~VectorTestCase() {}

VectorTestCase &VectorTestCase::operator=( VectorTestCase const & rhs )
{
	(void)rhs;
    return *this;
}

/*
** --------------------------------- METHODS ----------------------------------
*/
//Configure init function so it fills allTests var properly with your test funcitons
//and funciton names
void VectorTestCase::initTests()
{
    TestNameAndFunc t1 = 
    { 
        "testPushBack_PushBackOneElementToEmptyVector_SizeIsOneCapacityIsOneElementAddedCanBeRetreivedWithIndexZero", 
        testPushBack_PushBackOneElementToEmptyVector_SizeIsOneCapacityIsOneElementAddedCanBeRetreivedWithIndexZero
    };
    VectorTestCase::allTests.push_back(t1);
    TestNameAndFunc t2 = 
    { 
        "testPushBack_PushBackTwoElementsToEmptyVEctor_SizeIsTwoCapacityIsTwoElementAddedLastIsAccessibleAtIndexOne", 
        testPushBack_PushBackTwoElementsToEmptyVEctor_SizeIsTwoCapacityIsTwoElementAddedLastIsAccessibleAtIndexOne
    };
    VectorTestCase::allTests.push_back(t2);
    TestNameAndFunc t3 = 
    { 
        "testPushBack_PushBackThreeElementsToEmptyVector_SizeIsThreeCapacityIsFourElementAddedLastIsAccessibleAtIndexTwo", 
        testPushBack_PushBackThreeElementsToEmptyVector_SizeIsThreeCapacityIsFourElementAddedLastIsAccessibleAtIndexTwo
    };
    VectorTestCase::allTests.push_back(t3);
}

void VectorTestCase::run()
{
    initTests();
    startTestLoop();
}

/*
** --------------------------------- TEST_FUNCTIONS ----------------------------
*/
//write here definitions of test functions
//after fucnction proves to work correctly, change "CONSTRUCTION" to "DONE"
//RED
void VectorTestCase::testPushBack_PushBackOneElementToEmptyVector_SizeIsOneCapacityIsOneElementAddedCanBeRetreivedWithIndexZero()
{
    //Arrange
    ft::vector<int> actual;
    //Act
    actual.push_back(1);
    //Assert
    assertTrue(actual[0] == 1);
    assertTrue(actual.capacity() == 1);
    assertTrue(actual.size() == 1);
}
void VectorTestCase::testPushBack_PushBackTwoElementsToEmptyVEctor_SizeIsTwoCapacityIsTwoElementAddedLastIsAccessibleAtIndexOne()
{
    //Arrange
    ft::vector<int> actual;
    //Act
    actual.push_back(1);
    actual.push_back(2);
    //Assert
    assertTrue(actual.capacity() == 2);
    assertTrue(actual.size() == 2);
    assertTrue(actual[1] == 2);
}
void VectorTestCase::testPushBack_PushBackThreeElementsToEmptyVector_SizeIsThreeCapacityIsFourElementAddedLastIsAccessibleAtIndexTwo()
{
    //Arrange
    ft::vector<int> actual;
    //Act
    actual.push_back(1);
    actual.push_back(2);
    actual.push_back(3);
    //Assert
    assertTrue(actual.size() == 3);
    assertTrue(actual.capacity() == 4);
    assertTrue(actual[2] == 3);
}

/* ************************************************************************** */
