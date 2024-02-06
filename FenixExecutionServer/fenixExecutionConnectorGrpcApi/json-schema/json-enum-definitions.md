## enum LogPostStatusEnum
Each log post row is classified with the following
* INFO = 1; // The Log-post row is of Information Type
* WARNING = 2; // The Log-post row is of Warning Type, to inform user that something is not like it should
* VALIDATION_OK = 3; // The Log-post row is of Validation OK Type and should be used when the validation of the TestExecution matched expected result
* VALIDATION_ERROR = 4; // The Log-post row is of Validation Error Type and should be used when the validation of the TestExecution didn't matched expected result
* EXECUTION_OK = 5; // An execution, non validation, that was processed as expected
* EXECUTION_ERROR = 6; // An execution, non validation, that was not process as expected

## enum TestInstructionExecutionStatusEnum
Execution Enum for a TestInstruction Execution
* TIE_INITIATED = 1; // All set up for execution, but has not been triggered to start execution
* TIE_EXECUTING = 2; // TestInstruction is execution
* TIE_CONTROLLED_INTERRUPTION = 3; // Interrupted by in a controlled way
* TIE_CONTROLLED_INTERRUPTION_CAN_BE_RERUN = 4; // Interrupted by in a controlled way, but can be rerun
* TIE_FINISHED_OK = 5; // Finish as expected to TestInstruction definition
* TIE_FINISHED_OK_CAN_BE_RERUN = 6; // Finish as expected to TestInstruction definition, but can be rerun
* TIE_FINISHED_NOT_OK = 7; // Finish with errors in validations
* TIE_FINISHED_NOT_OK_CAN_BE_RERUN = 8; // Finish with errors in validations, but can be rerun
* TIE_UNEXPECTED_INTERRUPTION = 9; // The TestInstruction stopped executed in an unexpected way
* TIE_UNEXPECTED_INTERRUPTION_CAN_BE_RERUN = 10; // The TestInstruction stopped executed in an unexpected way, but can be rerun
* TIE_TIMEOUT_INTERRUPTION = 11; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}
* TIE_TIMEOUT_INTERRUPTION_CAN_BE_RERUN = 12; // The TestInstruction had a forced stop because of timeout due to {time.Now() > 'ExpectedExecutionEndTimeStamp'}, but can be rerun
