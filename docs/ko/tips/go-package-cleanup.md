# Go 패키지 정리 가이드 (Package Cleanup Guide)

> **목적**: Go 프로젝트에서 불필요한 패키지를 제거하고 의존성을 정리하는 표준 절차를 안내합니다.

## 핵심 요약 (Summary)

Go 의존성 관리는 **"코드 우선(Code-First)"**이 원칙입니다.
`go.mod`를 직접 수정하기보다, **소스 코드에서 사용을 중단하면 도구가 알아서 정리**해주는 방식을 따릅니다.

---

## 2. 표준 삭제 절차 (Standard Procedure)

### 1단계: 소스 코드 정리 (Clean Codebase)
프로젝트 내에서 해당 패키지를 사용하는 모든 코드를 제거합니다.

- `import` 구문 삭제
- 해당 패키지를 사용하는 함수/로직 삭제
- **중요**: 코드 생성 도구로 만들어진 파일(Generated Files)이 있다면 반드시 함께 삭제해야 합니다.

### 2단계: 의존성 정리 (Tidy Dependencies)
터미널에서 다음 명령어를 실행합니다.

```bash
go mod tidy
```

**이 명령어는 `go.mod`와 `go.sum`을 자동으로 관리합니다.**
1.  프로젝트의 모든 소스 코드를 스캔합니다.
2.  사용되지 않는 패키지(`require` 항목)를 `go.mod`에서 제거합니다.
3.  직접 의존성뿐만 아니라, 그 패키지가 가져왔던 **하위 의존성(Indirect Dependencies)**들도 함께 삭제합니다.
4.  패키지 무결성을 검증하는 `go.sum` 파일도 변경된 의존성에 맞춰 자동으로 갱신됩니다.

### 3단계 (선택): 강제 제거 (Force Remove)
`tidy` 실행 후에도 `go.mod`에 의존성이 남아있다면(예: 툴 의존성), 강제로 제거할 수 있습니다.

```bash
go mod edit -droprequire=module/path
go mod tidy
```

---

## 3. 예시: OpenAPI 제거 케이스 (Case Study)

> **상황**: `oapi-codegen` 및 관련 패키지를 제거하고 싶음.

### 잘못된 접근 ❌
- `go.mod`에서 `github.com/oapi-codegen/oapi-codegen/v2` 라인만 지움.
- 결과: `go build`나 `go mod tidy` 실행 시, 소스 코드에 import가 남아있어 다시 다운로드됨.

### 올바른 접근 ✅

1.  **파일 삭제**:
    - 생성된 코드 삭제: `internal/api/server.gen.go`, `internal/types/types.gen.go`
    - 스펙 파일 삭제: `api/openapi.yaml`
    - 설정 파일 삭제: `oapi-codegen.yaml`
2.  **코드 수정**:
    - `main.go`에서 `internal/api`를 import 하던 부분 삭제.
3.  **명령어 실행**:
    - `rm` 명령어로 파일들 제거 후
    - `go mod tidy` 실행 -> `oapi-codegen`과 하위 의존성(`spf13`, `mattn/go-runewidth` 등)이 자동으로 사라짐.

---

## 4. FAQ


**Q. 글로벌 저장소(`$GOPATH/pkg/mod`)에서도 지워지나요?**
A. 아니요. `go mod tidy`는 현재 프로젝트의 의존성만 끊습니다. 글로벌 캐시는 유지됩니다.
다른 프로젝트에서 공용으로 사용할 수 있기 때문입니다.

> **Tip: 글로벌 캐시 정리 방법**
> 디스크 용량 확보를 위해 전체 캐시를 비우고 싶다면 다음 명령어를 사용하세요:
> ```bash
> go clean -modcache
> ```

