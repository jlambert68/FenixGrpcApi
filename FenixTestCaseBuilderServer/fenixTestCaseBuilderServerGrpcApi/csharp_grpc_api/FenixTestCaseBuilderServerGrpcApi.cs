// <auto-generated>
//     Generated by the protocol buffer compiler.  DO NOT EDIT!
//     source: FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/fenixTestCaseBuilderServerGrpcApi.proto
// </auto-generated>
#pragma warning disable 1591, 0612, 3021
#region Designer generated code

using pb = global::Google.Protobuf;
using pbc = global::Google.Protobuf.Collections;
using pbr = global::Google.Protobuf.Reflection;
using scg = global::System.Collections.Generic;
namespace SubCustodyServer {

  /// <summary>Holder for reflection information generated from FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/fenixTestCaseBuilderServerGrpcApi.proto</summary>
  public static partial class FenixTestCaseBuilderServerGrpcApiReflection {

    #region Descriptor
    /// <summary>File descriptor for FenixTestCaseBuilderServer/fenixTestCaseBuilderServerGrpcApi/fenixTestCaseBuilderServerGrpcApi.proto</summary>
    public static pbr::FileDescriptor Descriptor {
      get { return descriptor; }
    }
    private static pbr::FileDescriptor descriptor;

    static FenixTestCaseBuilderServerGrpcApiReflection() {
      byte[] descriptorData = global::System.Convert.FromBase64String(
          string.Concat(
            "CmRGZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlci9mZW5peFRlc3RDYXNlQnVp",
            "bGRlclNlcnZlckdycGNBcGkvZmVuaXhUZXN0Q2FzZUJ1aWxkZXJTZXJ2ZXJH",
            "cnBjQXBpLnByb3RvEiFmZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNB",
            "cGkadkZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyL2Zlbml4VGVzdENhc2VC",
            "dWlsZGVyU2VydmVyR3JwY0FwaS9mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZl",
            "ckdycGNBcGlfQXZhaWxhYmxlTWVzc2FnZXMucHJvdG8afEZlbml4VGVzdENh",
            "c2VCdWlsZGVyU2VydmVyL2Zlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3Jw",
            "Y0FwaS9mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGlfR2VuZXJh",
            "bE1lc3NhZ2VzQW5kRW51bXMucHJvdG8adEZlbml4VGVzdENhc2VCdWlsZGVy",
            "U2VydmVyL2Zlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS9mZW5p",
            "eFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGlfQm9uZGluZ01lc3NhZ2Vz",
            "LnByb3RvGnVGZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlci9mZW5peFRlc3RD",
            "YXNlQnVpbGRlclNlcnZlckdycGNBcGkvZmVuaXhUZXN0Q2FzZUJ1aWxkZXJT",
            "ZXJ2ZXJHcnBjQXBpX1Rlc3RDYXNlTWVzc2FnZXMucHJvdG8afEZlbml4VGVz",
            "dENhc2VCdWlsZGVyU2VydmVyL2Zlbml4VGVzdENhc2VCdWlsZGVyU2VydmVy",
            "R3JwY0FwaS9mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGlfVGVz",
            "dEluc3RydWN0aW9uTWVzc2FnZXMucHJvdG8ahQFGZW5peFRlc3RDYXNlQnVp",
            "bGRlclNlcnZlci9mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkv",
            "ZmVuaXhUZXN0Q2FzZUJ1aWxkZXJTZXJ2ZXJHcnBjQXBpX1Rlc3RJbnN0cnVj",
            "dGlvbkNvbnRhaW5lck1lc3NhZ2VzLnByb3RvMssRCiZGZW5peFRlc3RDYXNl",
            "QnVpbGRlclNlcnZlckdycGNTZXJ2aWNlcxJ2CgtBcmVZb3VBbGl2ZRIxLmZl",
            "bml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5FbXB0eVBhcmFtZXRl",
            "choyLmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5BY2tOYWNr",
            "UmVzcG9uc2UiABLxAQo8TGlzdEFsbEF2YWlsYWJsZVRlc3RJbnN0cnVjdGlv",
            "bnNBbmRUZXN0SW5zdHJ1Y3Rpb25Db250YWluZXJzEjwuZmVuaXhUZXN0Q2Fz",
            "ZUJ1aWxkZXJTZXJ2ZXJHcnBjQXBpLlVzZXJJZGVudGlmaWNhdGlvbk1lc3Nh",
            "Z2UacS5mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuQXZhaWxh",
            "YmxlVGVzdEluc3RydWN0aW9uc0FuZFByZUNyZWF0ZWRUZXN0SW5zdHJ1Y3Rp",
            "b25Db250YWluZXJzUmVzcG9uc2VNZXNzYWdlIgAS/QEKQkxpc3RBbGxBdmFp",
            "bGFibGVQaW5uZWRUZXN0SW5zdHJ1Y3Rpb25zQW5kVGVzdEluc3RydWN0aW9u",
            "Q29udGFpbmVycxI8LmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0Fw",
            "aS5Vc2VySWRlbnRpZmljYXRpb25NZXNzYWdlGncuZmVuaXhUZXN0Q2FzZUJ1",
            "aWxkZXJTZXJ2ZXJHcnBjQXBpLkF2YWlsYWJsZVBpbm5lZFRlc3RJbnN0cnVj",
            "dGlvbnNBbmRQcmVDcmVhdGVkVGVzdEluc3RydWN0aW9uQ29udGFpbmVyc1Jl",
            "c3BvbnNlTWVzc2FnZSIAEpABChVMaXN0QWxsQXZhaWxhYmxlQm9uZHMSPC5m",
            "ZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuVXNlcklkZW50aWZp",
            "Y2F0aW9uTWVzc2FnZRo3LmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3Jw",
            "Y0FwaS5JbW1hdHVyZUJvbmRzTWVzc2FnZSIAEt0BCjlTYXZlQWxsUGlubmVk",
            "VGVzdEluc3RydWN0aW9uc0FuZFRlc3RJbnN0cnVjdGlvbkNvbnRhaW5lcnMS",
            "ai5mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuU2F2ZVBpbm5l",
            "ZFRlc3RJbnN0cnVjdGlvbnNBbmRQcmVDcmVhdGVkVGVzdEluc3RydWN0aW9u",
            "Q29udGFpbmVyc01lc3NhZ2UaMi5mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZl",
            "ckdycGNBcGkuQWNrTmFja1Jlc3BvbnNlIgASlwEKEExpc3RBbGxUZXN0Q2Fz",
            "ZXMSPi5mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuTGlzdFRl",
            "c3RDYXNlc1JlcXVlc3RNZXNzYWdlGj8uZmVuaXhUZXN0Q2FzZUJ1aWxkZXJT",
            "ZXJ2ZXJHcnBjQXBpLkxpc3RUZXN0Q2FzZXNSZXNwb25zZU1lc3NhZ2UiADAB",
            "EokBChNHZXREZXRhaWxlZFRlc3RDYXNlEjwuZmVuaXhUZXN0Q2FzZUJ1aWxk",
            "ZXJTZXJ2ZXJHcnBjQXBpLkdldFRlc3RDYXNlUmVxdWVzdE1lc3NhZ2UaMi5m",
            "ZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuVGVzdENhc2VNZXNz",
            "YWdlIgASwgEKH0xpc3RBbGxUZXN0Q2FzZVRlc3RJbnN0cnVjdGlvbnMSWy5m",
            "ZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuTGlzdEFsbFRlc3RJ",
            "bnN0cnVjdGlvbnNGb3JTcGVjaWZpY1Rlc3RDYXNlUmVxdWVzdE1lc3NhZ2Ua",
            "QC5mZW5peFRlc3RDYXNlQnVpbGRlclNlcnZlckdycGNBcGkuTWF0dXJlVGVz",
            "dEluc3RydWN0aW9uc01lc3NhZ2UiABLcAQooTGlzdEFsbFRlc3RDYXNlVGVz",
            "dEluc3RydWN0aW9uQ29udGFpbmVycxJkLmZlbml4VGVzdENhc2VCdWlsZGVy",
            "U2VydmVyR3JwY0FwaS5MaXN0QWxsVGVzdEluc3RydWN0aW9uQ29udGFpbmVy",
            "c0ZvclNwZWNpZmljVGVzdENhc2VSZXF1ZXN0TWVzc2FnZRpILmZlbml4VGVz",
            "dENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5NYXR1cmVUZXN0SW5zdHJ1Y3Rp",
            "b25Db250YWluZXJNZXNzYWdlIgASeAoMU2F2ZVRlc3RDYXNlEjIuZmVuaXhU",
            "ZXN0Q2FzZUJ1aWxkZXJTZXJ2ZXJHcnBjQXBpLlRlc3RDYXNlTWVzc2FnZRoy",
            "LmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5BY2tOYWNrUmVz",
            "cG9uc2UiABK0AQofU2F2ZUFsbFRlc3RDYXNlVGVzdEluc3RydWN0aW9ucxJb",
            "LmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5TYXZlQWxsVGVz",
            "dEluc3RydWN0aW9uc0ZvclNwZWNpZmljVGVzdENhc2VSZXF1ZXN0TWVzc2Fn",
            "ZRoyLmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0FwaS5BY2tOYWNr",
            "UmVzcG9uc2UiABLGAQooU2F2ZUFsbFRlc3RDYXNlVGVzdEluc3RydWN0aW9u",
            "Q29udGFpbmVycxJkLmZlbml4VGVzdENhc2VCdWlsZGVyU2VydmVyR3JwY0Fw",
            "aS5TYXZlQWxsVGVzdEluc3RydWN0aW9uQ29udGFpbmVyc0ZvclNwZWNpZmlj",
            "VGVzdENhc2VSZXF1ZXN0TWVzc2FnZRoyLmZlbml4VGVzdENhc2VCdWlsZGVy",
            "U2VydmVyR3JwY0FwaS5BY2tOYWNrUmVzcG9uc2UiAEIiWg0uL2dvX2dycGNf",
            "YXBpqgIQU3ViQ3VzdG9keVNlcnZlcmIGcHJvdG8z"));
      descriptor = pbr::FileDescriptor.FromGeneratedCode(descriptorData,
          new pbr::FileDescriptor[] { global::SubCustodyServer.FenixTestCaseBuilderServerGrpcApiAvailableMessagesReflection.Descriptor, global::SubCustodyServer.FenixTestCaseBuilderServerGrpcApiGeneralMessagesAndEnumsReflection.Descriptor, global::SubCustodyServer.FenixTestCaseBuilderServerGrpcApiBondingMessagesReflection.Descriptor, global::SubCustodyServer.FenixTestCaseBuilderServerGrpcApiTestCaseMessagesReflection.Descriptor, global::FenixTestCaseBuilderServerGrpcApi.FenixTestCaseBuilderServerGrpcApiTestInstructionMessagesReflection.Descriptor, global::SubCustodyServer.FenixTestCaseBuilderServerGrpcApiTestInstructionContainerMessagesReflection.Descriptor, },
          new pbr::GeneratedClrTypeInfo(null, null, null));
    }
    #endregion

  }
}

#endregion Designer generated code
